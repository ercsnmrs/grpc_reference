package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlog(c pb.BlogServiceClient) {
	log.Println("---ListBlog was invoked---")
	stream, err := c.ListBlogs(context.TODO(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Samething happened: %v\n", err)
			break
		}

		log.Println(res)
	}
}
