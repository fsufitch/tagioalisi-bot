package grpc

import (
	"fmt"
	"net"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
	"google.golang.org/grpc"

	pb "github.com/fsufitch/tagioalisi-bot/proto"
)

// TagioalisiGRPCServer is a struct describing a Tagioalisi gRPC interface
type TagioalisiGRPCServer struct {
	Port          config.GRPCPort
	Log           *log.Logger
	GreeterServer GreeterServer
}

// Run is a blocking function that starts and serves the gRPC server
func (s TagioalisiGRPCServer) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		s.Log.Errorf("failed to listen: %v", err)
		return fmt.Errorf("failed to listen: %w", err)
	}

	serv := grpc.NewServer()

	pb.RegisterGreeterServer(serv, &s.GreeterServer)
	s.Log.Infof("grpc server listening at %v", lis.Addr())
	if err := serv.Serve(lis); err != nil {
		s.Log.Errorf("failed to serve: %v", err)
		return fmt.Errorf("failed to serve: %w", err)
	}
	return fmt.Errorf("grpc server shut down with no error")
}
