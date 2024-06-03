genproto:
	protoc --go_out=. --go-grpc_out=. api/*.proto

start_server:
	go run cmd/grpc_server/main.go

start_streaming_server:
	go run cmd/grpc_streaming_server/main.go

start_client:
	go run cmd/grpc_client/main.go

start_streaming_client:
	go run cmd/grpc_streaming_client/main.go