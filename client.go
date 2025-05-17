package main

import (
	"context"
	"log"
	"time"

	"grpc-proto-golang/userpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect:", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.UserRequest{UserId: "889"}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling GetUser: %v", err)
	}

	log.Printf("Response: Name=%s, Age=%d, Date=%d", res.GetName(), res.GetAge(), res.GetDate())
}
