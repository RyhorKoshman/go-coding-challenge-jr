package main

import (
	"challenge/pkg/proto"
	"challenge/pkg/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const port = ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterChallengeServiceServer(s, &server.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
