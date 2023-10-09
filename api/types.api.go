package api

import "my-projects/stockbit/internal/usecase"

// Service to handle API requests
type Service struct {
	movie usecase.Movie
}
