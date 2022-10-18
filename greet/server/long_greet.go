package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	// For easy debugging
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving Req: %v\n", req)

		res += fmt.Sprintf("Hello %s! ", req.FirstName)
	}
}
