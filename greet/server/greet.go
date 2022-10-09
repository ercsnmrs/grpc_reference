// Will contain the Greet RPC Endpoint
package main

import (
	"context"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	// For easy debugging
	log.Printf("Greet function was invoked with %v\n", req)

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
