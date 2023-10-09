package main

import (
	"log"

	gops "github.com/google/gops/agent"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	googleGRPC "google.golang.org/grpc"
)

func main() {
	// Initialize everything here
	grpcServer := initializeService()

	serve(grpcServer)

}

func initializeService() *googleGRPC.Server {

	// GOPS
	if err := gops.Listen(gops.Options{
		ShutdownCleanup: true,
	}); err != nil {
		log.Fatal("func initializeService Can't initialize GOPS", err)
	}

	// Initialize GRPC Server
	grpcServer := initGRPC()

	return grpcServer
}
