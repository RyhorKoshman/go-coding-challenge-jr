package main

import (
	"challenge/pkg/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewChallengeServiceClient(conn)
	md := metadata.Pairs("i-am-random-key", "123")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	url := "https://example.com"
	data, err := c.MakeShortLink(ctx, &proto.Link{Data: url})
	fmt.Println(data)
	fmt.Println(err)
}
