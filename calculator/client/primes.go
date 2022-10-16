package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) error {
	// For easy debugging
	log.Printf("doPrimes function was invoked")

	req := &pb.PrimesRequest{
		Number: 210,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("doPrimes response: %v\n", res.Result)
	}

	return nil
}
