package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-proto-golang/userpb"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Printf("Received request for user_id: %s\n", req.GetUserId())
	return &userpb.UserResponse{
		Name: "Chandra Febrian",
		Age:  20,
		Date: 20231010,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
