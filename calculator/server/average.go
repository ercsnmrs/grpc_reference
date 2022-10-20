// Will contain the Greet RPC Endpoint
package main

import (
	"io"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	// For easy debugging
	log.Printf("Average function was invoked")
	
	sum := 0
	counter := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(sum) / float64(counter),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving Req: %v\n", req)
		counter++
		sum += int(req.Number)
	}
}
