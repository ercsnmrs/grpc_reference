package main

import (
	"context"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

func doSum(c pb.CalulatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Addends1: 3,
		Addends2: 10,
	})

	if err != nil {
		log.Fatalf("Could not add: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}
