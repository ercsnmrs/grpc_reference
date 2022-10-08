package main

import (
	"log"
	"net"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

// Create Grpc Server Object
// We will be able to implement all the RPC endpoints that we will define in our greet proto
type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
