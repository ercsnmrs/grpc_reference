package main

import (
	"log"
	"net"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
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

	opts := []grpc.ServerOption{}
	tls := false //change that to false if needed
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading cerificates: %v\n", err)
		}

		// Add this credentials to a ServerOption.
		// What we can do is that we are going to define a options variable
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	// Register Service Server
	// Use the service to implement the rpc endpoints
	pb.RegisterGreetServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
