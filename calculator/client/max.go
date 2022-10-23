package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("doMax function was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	// reqs := []*pb.MaxRequest{
	// 	{ Number: 15},
	// 	{ Number: 32},
	// 	{ Number: 10},
	// 	{ Number: 6},
	// 	{ Number: 100},
	// }

	go func() {
		numbers := []int32{15, 32, 10, 6, 100}
		for _, number := range numbers {
			log.Printf("Sending Request: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
