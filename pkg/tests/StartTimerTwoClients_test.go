package main

import (
	"challenge/pkg/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestTwo(t *testing.T) {
	const TimeName = "sldfnsdnkl"
	const Seconds = 12
	const Frequency = 3

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	conn2, err2 := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	defer conn2.Close()
	client := proto.NewChallengeServiceClient(conn)
	client2 := proto.NewChallengeServiceClient(conn2)
	timerRequest := &proto.Timer{
		Seconds:   Seconds,
		Name:      TimeName,
		Frequency: Frequency,
	}
	stream1, err := client.StartTimer(context.Background(), timerRequest)
	stream2, err := client2.StartTimer(context.Background(), timerRequest)
	if err != nil {
		log.Fatalf("Error calling StartTimer: %v", err)
	}
	T := int64(Seconds)
	var total = 0
	for {
		_, e := stream1.Recv()
		total++
		T -= Frequency
		if e != nil {
			break
		}
	}
	T = int64(Seconds)

	for {
		_, e := stream2.Recv()
		total--
		T -= Frequency
		if e != nil {
			break
		}
	}
	if total != 0 {
		t.Errorf("Different number of messages")
	}
}
