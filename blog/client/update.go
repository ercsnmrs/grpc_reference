package main

import (
	"context"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/blog/proto"
)

func UpdateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---UpdateBlog was invoked---")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Mr. Peanuts",
		Title:    "My Second RPC Blog",
		Content:  "Content of the second blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog has been updated!")
}
