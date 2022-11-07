package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ercsnmrs/grpc_reference/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	ReadBlog(c, id) // should be valid
	// ReadBlog(c, "InvaliDIDHere")
	UpdateBlog(c, id)
	ListBlog(c)
}
