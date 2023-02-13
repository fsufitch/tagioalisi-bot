package grpc

import (
	"github.com/fsufitch/tagioalisi-bot/proto"
	"github.com/google/wire"
)

// GRPCProviderSet provides everything needed run the gRPC server
var GRPCProviderSet = wire.NewSet(
	wire.Struct(new(GreeterServer), "*"),
	wire.Struct(new(proto.UnimplementedGreeterServer), "*"),

	ProvideTagioalisiGRPC,
	ProvideWrapGRPCWebsocketFunc,
)
