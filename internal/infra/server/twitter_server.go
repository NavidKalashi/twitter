package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/NavidKalashi/twitter/internal/adapters/http/grpcTwitter"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTwitterServiceServer
	mu sync.Mutex
	tweets []*pb.Tweet
}

func (s *server) CreateTweet(ctx context.Context, req *pb.CreateTweetRequest) (*pb.CreateTweetResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tweet := &pb.Tweet{
		Id: int32(len(s.tweets) + 1),
		User: req.User,
		Text: req.Text,
		Timestamp: time.Now().Unix(),
	}
	s.tweets = append(s.tweets, tweet)
	return &pb.CreateTweetResponse{Tweet: tweet}, nil
}

func (s *server) GetTweet(ctx context.Context, req *pb.GetTweetRequest) (*pb.GetTweetResponse, error){
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, tweet := range s.tweets {
		if tweet.Id == req.Id {
			return &pb.GetTweetResponse{Tweet: tweet}, nil
		}
	}
	return nil, fmt.Errorf("tweet not found with id: %d", req.Id)
}

func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTwitterServiceServer(s, &server{})
	log.Println("Server is running on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}