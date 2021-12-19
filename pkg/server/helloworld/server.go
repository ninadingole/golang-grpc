package helloworld

import (
	"context"
)

type Server struct {
	UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + req.GetName()}, nil
}
