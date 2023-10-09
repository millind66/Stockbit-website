package service

import (
	"my-projects/stockbit/internal/entity"
	"my-projects/stockbit/internal/repository"
)

type movieSVC struct {
	movieRP    repository.Movie
	databaseRP repository.Database
}

// Movie service handles movie related business logic
type Movie interface {
	Search(request entity.SearchRequest) (entity.SearchResponse, error)
	Get(request entity.GetRequest) (entity.GetResponse, error)
}
