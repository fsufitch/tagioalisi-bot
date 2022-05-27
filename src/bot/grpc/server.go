package grpc

import (
	"net/http"

	"github.com/fsufitch/tagioalisi-bot/log"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	pb "github.com/fsufitch/tagioalisi-bot/proto"
)

type TagioalisiGRPC *grpc.Server

func ProvideTagioalisiGRPC(
	log *log.Logger,
	// Add other GRPC services here, as they become supported
	GreeterServer *GreeterServer,
) TagioalisiGRPC {
	serv := grpc.NewServer()

	// Add other GRPC services here, as they become supported
	pb.RegisterGreeterServer(serv, GreeterServer)

	return TagioalisiGRPC(serv)
}

type WrapGRPCWebsocketFunc func(inner http.Handler) http.Handler

func ProvideWrapGRPCWebsocketFunc(serv TagioalisiGRPC) WrapGRPCWebsocketFunc {
	webserv := grpcweb.WrapServer(serv,
		grpcweb.WithOriginFunc(func(string) bool { return true }),
		grpcweb.WithAllowNonRootResource(true),
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(*http.Request) bool { return true }), // TODO: Tighten this?
	)

	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if webserv.IsGrpcWebSocketRequest(r) {
				webserv.HandleGrpcWebsocketRequest(w, r)
				return
			}
			inner.ServeHTTP(w, r)
		})
	}
}
