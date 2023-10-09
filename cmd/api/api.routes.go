package main

import (
	"my-projects/stockbit/api"
	"net/http"
)

// Initialize all routes for API
func initRoutes(srv *api.Service) {
	http.HandleFunc("/search", srv.SearchMovie)
	http.HandleFunc("/get", srv.GetMovie)
}
