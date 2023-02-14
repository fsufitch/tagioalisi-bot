package grpc

import (
	"fmt"
	"net"

	"github.com/fsufitch/tagioalisi-bot/config"
	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/fsufitch/tagioalisi-bot/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type TagioalisiGRPC struct {
	Log           *log.Logger
	Port          config.GRPCPort
	GreeterServer GreeterServer
}

func (tgrpc TagioalisiGRPC) Run() error {
	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServer, tgrpc.GreeterServer)
	addr := fmt.Sprintf(":%d", tgrpc.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "could not open port for grpc")
	}

	tgrpc.Log.Infof("grpc: starting server on addr: %s", addr)

	return grpcServer.Serve(lis)
}
