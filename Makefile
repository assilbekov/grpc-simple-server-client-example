genproto:
	protoc --go_out=. --go-grpc_out=. api/example.proto

start_server:
	go run cmd/grpc_server/main.go

start_client:
	go run cmd/grpc_client/main.go