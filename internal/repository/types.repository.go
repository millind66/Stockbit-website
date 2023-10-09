package repository

import (
	"my-projects/stockbit/internal/entity"
	"net/http"
)

type movieRP struct {
	httpClient *http.Client
}

type databaseRP struct {
	dbClient entity.DBClient
}

// Movie repository handles movie related third party integration
type Movie interface {
	Search(request entity.SearchRequest) (entity.SearchResponse, error)
	Get(request entity.GetRequest) (entity.GetResponse, error)
}

// Database repository handles database related transaction
type Database interface {
	InsertLog(data entity.Log) error
}
