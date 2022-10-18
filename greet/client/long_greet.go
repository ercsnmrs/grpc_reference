package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	// For easy debugging
	log.Printf("doLongGreet function was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Ericson"},
		{FirstName: "Stephy"},
		{FirstName: "Sunni"},
		{FirstName: "Peanuts"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
