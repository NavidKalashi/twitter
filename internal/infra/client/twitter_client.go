package client

import (
	"context"
	"log"
	"time"

	pb "github.com/NavidKalashi/twitter/internal/adapters/http/grpcTwitter"
	"google.golang.org/grpc"
)


func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTwitterServiceClient(conn)

	// Create a tweet
	user := &pb.User{Id: 1, Username: "user1", Name: "User one"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	createTweetResp, err := client.CreateTweet(ctx, &pb.CreateTweetRequest{User: user, Text: "hello world"})
	if err != nil {
		log.Fatalf("could not create tweet: %v", err)
	}
	log.Printf("Created tweet: %v", createTweetResp.Tweet)

	// Get the tweet
	getTweetResp, err := client.GetTweet(ctx, &pb.GetTweetRequest{Id: createTweetResp.Tweet.Id})
	if err != nil {
		log.Fatalf("could not get tweet: %v", err)
	}
	log.Printf("retrieved tweet: %v", getTweetResp.Tweet)
}