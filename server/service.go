package server

import (
	"context"
	"test_grpc/pb"
)

func (s *Server) HelloWorld(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Pong: "Hello World"}, nil
}
