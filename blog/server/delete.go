package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ercsnmrs/grpc_reference/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, req *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v \n", req)

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot parse ID",
		)
	}

	filter := bson.M{"_id": oid}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}

	return &emptypb.Empty{}, nil
}
