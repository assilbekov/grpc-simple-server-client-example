genproto:
	protoc --go_out=. --go-grpc_out=. api/example.proto
