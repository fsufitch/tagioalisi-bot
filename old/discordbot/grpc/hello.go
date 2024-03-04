package grpc

import (
	"context"
	"fmt"

	"github.com/fsufitch/tagioalisi-bot/log"
	pb "github.com/fsufitch/tagioalisi-bot/proto"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
	Log *log.Logger
}

func (s GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	s.Log.Infof("Client says hello! name=%s", request.Name)
	reply := pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s! I am the server!", request.Name),
	}
	return &reply, nil
}
