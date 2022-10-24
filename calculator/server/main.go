package main

import (
	"log"
	"net"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

// Create Grpc Server Object
// We will be able to implement all the RPC endpoints that we will define in our greet proto
type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	// Register Service Server
	// Use the service to implement the rpc endpoints

	pb.RegisterCalculatorServiceServer(s, &Server{})
	// reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
