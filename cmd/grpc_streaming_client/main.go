package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-simple-server-client-example/api/proto/streaming_example"
	"log"
	"time"
)

func printFeature(client streaming_example.RouteGuideClient, point *streaming_example.Point) {
	fmt.Printf("Looking for feature at point (%d, %d)\n", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		log.Fatalf("%v.GetFeature(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}

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

	client := streaming_example.NewRouteGuideClient(conn)

	printFeature(client, &streaming_example.Point{Latitude: 409146138, Longitude: -746188906})
}
