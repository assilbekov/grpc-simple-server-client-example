package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-simple-server-client-example/api/proto/streaming_example"
	"io"
	"log"
	"math/rand"
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

func printFeatures(client streaming_example.RouteGuideClient, rect *streaming_example.Rectangle) {
	fmt.Printf("Looking for features within %v\n", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}

func randomPoint(r *rand.Rand) *streaming_example.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &streaming_example.Point{Latitude: lat, Longitude: long}
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

	// Looking for a valid feature
	printFeature(client, &streaming_example.Point{Latitude: 409146138, Longitude: -746188906})

	// Feature missing.
	printFeature(client, &streaming_example.Point{Latitude: 0, Longitude: 0})

	// Looking for features between 40, -75 and 42, -73.
	printFeatures(client, &streaming_example.Rectangle{
		Lo: &streaming_example.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &streaming_example.Point{Latitude: 420000000, Longitude: -730000000},
	})
}
