proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

path:
	export PATH=$PATH:~/go/bin	
	
main:
	go run main.go
	
.PHONY: proto main