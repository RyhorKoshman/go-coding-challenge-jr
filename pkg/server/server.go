package server

import "challenge/pkg/proto"

type Server struct {
	proto.UnimplementedChallengeServiceServer
}
