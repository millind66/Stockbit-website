package usecase

import (
	"my-projects/stockbit/internal/entity"
	"my-projects/stockbit/internal/service"
)

// NewMovie to create new movie usecase
func NewMovie(movieSVC service.Movie) Movie {
	return &movieUC{
		movieSVC,
	}
}

func (uc *movieUC) Search(request entity.SearchRequest) (entity.SearchResponse, error) {
	return uc.movieSVC.Search(request)
}

func (uc *movieUC) Get(request entity.GetRequest) (entity.GetResponse, error) {
	return uc.movieSVC.Get(request)
}
