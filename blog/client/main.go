package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Amartya-Bhardwaj/grpc/blog/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func CreateBlog(c pb.BlogServiceClient) string {
	log.Println("===CreateBlog===invoked")
	blog := &pb.Blog{
		AuthorId: "Amartya",
		Title:    "First",
		Content:  "Hello world",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Blog has been created %s\n", res.GetId())
	return res.GetId()
}


func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	CreateBlog(c)
}
