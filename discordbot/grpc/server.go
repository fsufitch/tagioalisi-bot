package grpc

import (
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
	// addr := fmt.Sprintf(":%d", tgrpc.Port)
	addr := ":9999"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "could not open port for grpc")
	}

	return grpcServer.Serve(lis)

	// var corsHandlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "*")

	// 	if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
	// 		fmt.Println("??????????")
	// 		grpcServer.ServeHTTP(w, r)
	// 	} else {
	// 		w.Header().Set("Content-Type", "text/plain")
	// 		w.WriteHeader(http.StatusOK)
	// 		fmt.Fprintf(w, "Sir, this is a GRPC server\n")
	// 		fmt.Fprintf(w, "HTTP version: %s\n", r.Proto)
	// 		fmt.Fprintf(w, "Content-Type: %s\n", r.Header.Get("Content-Type"))
	// 	}
	// 	fmt.Println(w.Header())
	// }

	// httpServer := http.Server{}
	// httpServer.Handler = corsHandlerFunc
	// http2.ConfigureServer(&httpServer, &http2.Server{})

	// addr := fmt.Sprintf(":%d", tgrpc.Port)
	// lis, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	return errors.Wrap(err, "could not open port for grpc")
	// }

	// certFile := os.Getenv("HTTPS_CERT")
	// keyFile := os.Getenv("HTTPS_KEY")
	// if certFile != "" && keyFile != "" {
	// 	tgrpc.Log.Infof("grpc: starting server on addr: %s (HTTPS)", addr)
	// 	return httpServer.ServeTLS(lis, certFile, keyFile)

	// } else {
	// 	tgrpc.Log.Infof("grpc: starting server on addr: %s (HTTP)", addr)
	// 	return httpServer.Serve(lis)
	// }

}
