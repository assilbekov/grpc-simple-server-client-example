package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-simple-server-client-example/api/proto/example"
	"log"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close: %v", err)
		}
	}(conn)

	client := example.NewExampleServiceClient(conn)
	resp, err := client.ExampleMethod(context.Background(), &example.ExampleRequest{
		ExampleField: "example request",
	})
	if err != nil {
		log.Fatalf("failed to call ExampleMethod: %v", err)
	}
	fmt.Println(resp.ExampleField)
}
