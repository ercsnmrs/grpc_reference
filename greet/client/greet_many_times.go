package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) error {
	// For easy debugging
	log.Printf("GreetManyTimes function was invoked")

	req := &pb.GreetRequest{
		FirstName: "Mr. Peanut",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}

	return nil
}
