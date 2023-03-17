package main

import (
	
	"database/sql"

	"fmt"
	"log"
	"net"
	"os"
	"sync"

	db "github.com/obasootom/gochatapp/db/sqlc"
	"github.com/obasootom/gochatapp/pb"
	"github.com/obasootom/gochatapp/util"

	_ "github.com/lib/pq"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

type Connection struct {
	stream pb.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

type Server struct {
	Connection []*Connection
	pb.UnimplementedBroadcastServer
	store *db.Store
}

func (s *Server) CreateStream(pconn *pb.Connect, stream pb.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *pb.Message) (*pb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)
	arg := db.CreatePostParams{
		Username: msg.Id,
		Content:  msg.Content,
	}
	post, err := s.store.CreatePost(ctx, arg)
	if err != nil {
		fmt.Printf("cannot send message %s", err)
	}
	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *pb.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Info("Sending message to: ", conn.stream, post)

				if err != nil {
					grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)
	}
	go func() {
		wait.Wait()
		close(done)
	}()
	<-done
	return &pb.Close{}, nil
}

func main() {
	config, err := util.NewConfigure(".")
	if err != nil {
		fmt.Println("cannot create config")
	}
	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		fmt.Println("cannot connect to database")
	}
	store := db.NewStore(conn)
	LoadGrpc(store)

}

func LoadGrpc(store *db.Store) {
	var connections []*Connection

	server := &Server{connections, pb.UnimplementedBroadcastServer{}, store}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	grpcLog.Info("Starting server at port :8888")

	pb.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)
}
