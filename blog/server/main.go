package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ercsnmrs/grpc_reference/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

// Create Grpc Server Object
// We will be able to implement all the RPC endpoints that we will define in our greet proto
type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	// Register Service Server
	// Use the service to implement the rpc endpoints

	pb.RegisterBlogServiceServer(s, &Server{})
	// reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
