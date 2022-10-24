package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/ercsnmrs/grpc_reference/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	// Create connection with grpc using Dial(addr) func

	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		// Add this credentials to a ServerOption.
		// What we can do is that we are going to define a options variable
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	// defer will execute XXX at the end of the function
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetWithDeadline
	// doGreetWithDeadline(c, 1*time.Second)
}
