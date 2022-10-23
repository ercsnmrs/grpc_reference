// Will contain the Greet RPC Endpoint
package main

import (
	"io"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var (
		// current int32
		maximum int32
	)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		// if current < req.Number {
		// 	current = req.Number
		// }

		// log.Printf("Current Highest Number: %v", current)

		// err = stream.Send(&pb.MaxResponse{
		// 	Result: current,
		// })

		if number := req.Number; number > maximum {
			maximum = number
			err = stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client %v\n", err)
			}
		}
	}
}
