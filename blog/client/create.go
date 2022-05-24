package main

// import (
// 	"context"
// 	"log"
	
// 	pb "github.com/Amartya-Bhardwaj/grpc/blog/proto"
// )

// func CreateBlog(c pb.BlogServiceClient) string {
// 	log.Println("===CreateBlog===invoked")
// 	blog := &pb.Blog{
// 		AuthorId: "Amartya",
// 		Title:    "First",
// 		Content:  "Hello world",
// 	}
// 	res, err := c.CreateBlog(context.Background(), blog)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Blog has been created %s\n", res.GetId())
// 	return res.GetId()
// }
