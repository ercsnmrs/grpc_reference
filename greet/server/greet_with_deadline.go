// Will contain the Greet RPC Endpoint
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline function was invoked with %v\n", req)

	// Goals - Implement forloop to check if deadline is exceeded
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceld the reuqest")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
