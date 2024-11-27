package main

import (
	"challenge/pkg/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"testing"
)

func getShortLink(url string) (*proto.Link, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewChallengeServiceClient(conn)
	md := metadata.New(map[string]string{})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return c.MakeShortLink(ctx, &proto.Link{Data: url})
}

func normalizeLink(s string) string {
	if s[len(s)-1] == '/' {
		return s[:len(s)-1]
	}
	return s
}

func TestCorrectnessCheck(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/Fast_Fourier_transform"
	data, err := getShortLink(url)

	if err != nil {
		t.Error(err)
	}
	shortUrl := data.Data
	resp, err := http.Get(shortUrl)
	if err != nil {
		t.Errorf("http.Get error: %v", err)
	}
	decUrl := resp.Request.URL.String()
	if normalizeLink(decUrl) != normalizeLink(url) {
		t.Error("Different links ", normalizeLink(shortUrl), " ", normalizeLink(url))
	}
}
