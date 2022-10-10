package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	// Create connection with grpc using Dial(addr) func
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	// defer will execute XXX at the end of the function
	defer conn.Close()

	c:= pb.NewCalulatorServiceClient(conn)

	doSum(c)
}
