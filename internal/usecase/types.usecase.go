package usecase

import (
	"my-projects/stockbit/internal/entity"
	"my-projects/stockbit/internal/service"
)

type movieUC struct {
	movieSVC service.Movie
}

// Movie usecase handles movie related endpoint
type Movie interface {
	Search(request entity.SearchRequest) (entity.SearchResponse, error)
	Get(request entity.GetRequest) (entity.GetResponse, error)
}
