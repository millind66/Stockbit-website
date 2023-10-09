package main

import (
	"log"
	"os"
	"os/signal"

	googleGRPC "google.golang.org/grpc"
)

func serve(grpcServer *googleGRPC.Server) {

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT

	// Shutdown grpc server
	if grpcServer != nil {
		grpcServer.Stop()
	}

	log.Println("func serve Gracfully Stopped : Hotel Discovery GRPC")

}
