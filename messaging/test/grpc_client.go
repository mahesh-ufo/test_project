package main

import (
	"log"
	"os"
	"time"

	pb "github.com/mahesh-ufo/test_project/messaging/schema"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50001"
	defaultName = "Request"
)

func main() {
	// Set up a connection to the server.
	var conn, err = grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	var c = pb.NewAdvtEngineMessagingClient(conn)

	// Contact the server and print out its response.
	var name = defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.TestRPC(ctx, &pb.Request{Name: name})

	if err != nil {
		log.Fatalf("could not call TestRPC: %v", err)
	}

	log.Printf("Response: %s", r.Message)
}
