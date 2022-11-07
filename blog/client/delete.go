package main

import (
	"context"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/blog/proto"
)

func DeleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---DeleteBlog was invoked---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted!")
}
