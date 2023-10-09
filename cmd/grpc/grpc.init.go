package main

import (
	"fmt"
	"log"
	"my-projects/stockbit/cmd/internal"

	"my-projects/stockbit/grpc"
	pb "my-projects/stockbit/proto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/valyala/fasthttp/reuseport"
	googleGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func initGRPC() *googleGRPC.Server {
	// Init server
	s := googleGRPC.NewServer(
		googleGRPC.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
		)),

		googleGRPC.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
		)),
	)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	usecases := internal.GetUsecase()
	grpcSrv := grpc.NewGRPCService(usecases.Movie)

	pb.RegisterStockbitServer(s, grpcSrv)

	// Start GRPC
	go func(srv *googleGRPC.Server) {
		l, err := reuseport.Listen("tcp4", ":9889")
		if err != nil {
			log.Fatal("func Start failed to start grpc server", err)
		}

		fmt.Println("Running stockbit-grpc service on :9889")
		err = srv.Serve(l)
		if err != nil {
			log.Fatal("func initGRPC Can't start GRPC", err)
		}
	}(s)

	return s
}
