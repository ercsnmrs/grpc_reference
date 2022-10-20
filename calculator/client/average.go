package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	// For easy debugging
	log.Println("doAverage function was invoked")

	// reqs := []*pb.AverageRequest{
	// 	{Number: 1},
	// 	{Number: 2},
	// 	{Number: 3},
	// 	{Number: 4},
	// }

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	numbers := []int32{1, 2, 3, 4}

	for _, req := range numbers {
		log.Printf("Sending req: %v\n", req)
		//stream.Send(req)
		stream.Send(&pb.AverageRequest{
			Number: req,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average Result: %v\n", res.Result)
}
