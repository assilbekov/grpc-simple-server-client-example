package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	"grpc-simple-server-client-example/api/proto/example"
)

type server struct {
	example.UnimplementedExampleServiceServer
}

func (s *server) ExampleMethod(
	ctx context.Context,
	req *example.ExampleRequest,
) (*example.ExampleResponse, error) {
	return &example.ExampleResponse{
		ExampleField: "example response: " + req.ExampleField,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	example.RegisterExampleServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
