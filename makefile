proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

path:
	export PATH=$PATH:~/go/bin	
	
main:
	go run main.go

minikube:
	eval $(minikube docker-env)

migrate:
	migrate create -ext sql -dir db/migration -seq init-schema
	
sqlc:
	sqlc generate

createdb:
	docker exec -it postgresline createdb --username=root --owner=root gochatapp

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:8000/gochatapp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:8000/gochatapp?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgresline dropdb gochatapp

.PHONY: proto main migrate sqlc createdb dropdb migrateup migratedown