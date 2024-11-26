package main

import (
	"challenge/pkg/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	const TimeName = "sldfnsdnkl"
	const Seconds = 12
	const Frequency = 3

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	client := proto.NewChallengeServiceClient(conn)
	timerRequest := &proto.Timer{
		Seconds:   Seconds,
		Name:      TimeName,
		Frequency: Frequency,
	}
	stream, err := client.StartTimer(context.Background(), timerRequest)

	if err != nil {
		log.Fatalf("Error calling StartTimer: %v", err)
	}
	T := int64(Seconds)
	for {
		TTT, e := stream.Recv()
		T -= Frequency
		if e != nil {
			break
		}
		if TTT.Seconds != T {
			t.Errorf("TTT.Seconds = %d, want %d", TTT.Seconds, T)
		}
		t.Log(TTT)
	}
}
