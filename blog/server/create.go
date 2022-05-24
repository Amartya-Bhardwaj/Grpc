package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Amartya-Bhardwaj/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBLog function invoked %v",in)
	data := BlogItem{
		AuthorId: in.AuthorId,
		Title: in.Title,
		Content: in.Content,
	}

	res,err:= collection.InsertOne(ctx,data)
	if err!=nil{
		return nil,status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error %v\n",err),
		)
	}

	oid,ok := res.InsertedID.(primitive.ObjectID)
	if !ok{
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	},nil
}
