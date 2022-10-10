// Will contain the Greet RPC Endpoint
package main

import (
	"context"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	// For easy debugging
	log.Printf("Sum function was invoked with %v\n", req)

	sum := req.Addends1 + req.Addends2

	return &pb.SumResponse{
		Result: sum,
	}, nil
}
