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
	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, tgrpc.GreeterServer)

	addr := fmt.Sprintf(":%d", tgrpc.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "could not open port for grpc")
	}
	tgrpc.Log.Infof("grpc: starting server on addr: %s", addr)

	return server.Serve(lis)
}

// type WrapGRPCWebsocketFunc func(inner http.Handler) http.Handler

// func ProvideWrapGRPCWebsocketFunc(serv TagioalisiGRPC) WrapGRPCWebsocketFunc {
// 	webserv := grpcweb.WrapServer(serv.server,
// 		grpcweb.WithOriginFunc(func(string) bool { return true }),
// 		grpcweb.WithAllowNonRootResource(true),
// 		grpcweb.WithWebsockets(true),
// 		grpcweb.WithWebsocketOriginFunc(func(*http.Request) bool { return true }), // TODO: Tighten this?
// 	)

// 	return func(inner http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// w.Header().Set("Access-Control-Allow-Origin", "*")
// 			if webserv.IsGrpcWebSocketRequest(r) {
// 				fmt.Printf("ws request %v\n", r)
// 				webserv.HandleGrpcWebsocketRequest(w, r)
// 				return
// 			}
// 			fmt.Printf("http request %v\n", r)
// 			inner.ServeHTTP(w, r)
// 		})
// 	}
// }
