package main

import (
	"log"
	"os"
	"time"

	pb "../protobuff"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50001"
	defaultName = "Request"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTestMessagingClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.TestRPC(ctx, &pb.TestRequest{Name: name})

	if err != nil {
		log.Fatalf("could not call TestRPC: %v", err)
	}

	log.Printf("Response: %s", r.Message)
}
