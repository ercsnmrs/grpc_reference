// Will contain the Greet RPC Endpoint
package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	// For easy debugging
	log.Printf("Sqrt function was invoked with %v\n", req)

	number := req.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
