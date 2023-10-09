package main

import (
	"fmt"
	"log"
	"my-projects/stockbit/api"
	"my-projects/stockbit/cmd/internal"
	"net/http"

	gops "github.com/google/gops/agent"
)

func main() {

	// GOPS
	if err := gops.Listen(gops.Options{
		ShutdownCleanup: true,
	}); err != nil {
		log.Fatal("[FATAL] Can't initialize GOPS", err)
	}

	usecases := internal.GetUsecase()

	// Init routes & services
	apiSrv := api.NewAPIService(usecases.Movie)
	initRoutes(apiSrv)

	// Set routes and serve HTTP
	fmt.Println("Running stockbit-api service on :9000")
	http.ListenAndServe(":9000", nil)
}
