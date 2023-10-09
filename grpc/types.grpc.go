package grpc

import "my-projects/stockbit/internal/usecase"

// Service to handle GRPC requests
type Service struct {
	movie usecase.Movie
}
