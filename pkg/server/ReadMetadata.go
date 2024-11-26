package server

import (
	"challenge/pkg/proto"
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func (*Server) ReadMetadata(ctx context.Context, in *proto.Placeholder) (*proto.Placeholder, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	seq := md.Get(in.Data)
	if len(seq) == 0 {
		return nil, errors.New("no sequence in metadata")
	}
	return &proto.Placeholder{Data: seq[0]}, nil
}
