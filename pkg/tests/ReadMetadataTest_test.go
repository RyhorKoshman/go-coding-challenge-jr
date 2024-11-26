package main

import (
	"challenge/pkg/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"testing"
)

func TestReadMetadata(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	S := "abpba"
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewChallengeServiceClient(conn)
	md := metadata.Pairs("i-am-random-key", S)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	c.ReadMetadata(ctx, &proto.Placeholder{Data: "i-am-random-key"})
	data, _ := c.ReadMetadata(ctx, &proto.Placeholder{Data: "i-am-random-key"})
	t.Log(data)
	if data.Data != S {
		t.Errorf("ReadMetadata returned %s instead of %s", data.Data, S)
	}
}
