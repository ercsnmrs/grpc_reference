package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	// For easy debugging
	log.Printf("doGreetEveryone function was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{ FirstName: "Ericson"},
		{ FirstName: "Stephy"},
		{ FirstName: "Peamnuts"},
	}

	// Create a go channel
	// Run two go routines
	// 1st Routine: Send all request
	// 2nd Routine: Receive responses from server

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send Request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
	
			if err != nil {
				log.Printf("Error while receiving : %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
