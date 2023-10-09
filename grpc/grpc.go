package grpc

import (
	"errors"
	"log"
	"my-projects/stockbit/internal/entity"
	"my-projects/stockbit/internal/usecase"
	pb "my-projects/stockbit/proto"

	context "golang.org/x/net/context"
)

// NewGRPCService to create new GRPC service
func NewGRPCService(movieUC usecase.Movie) *Service {
	return &Service{
		movie: movieUC,
	}
}

// SearchMovie handles GRPC requests for search movie
func (srv *Service) SearchMovie(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	resp := &pb.SearchResponse{}
	// validate request
	if req == nil || req.SearchWord == "" || req.Pagination < 1 {
		log.Println("func SearchMovie invalid request")
		return resp, errors.New("invalid request")
	}

	// search movie from usecase
	result, err := srv.movie.Search(entity.SearchRequest{
		SearchWord: req.SearchWord,
		Pagination: int(req.Pagination),
	})
	if err != nil {
		log.Println("func SearchMovie error from usecase", err)
		return resp, err
	}

	// parse result struct to proto
	resp.Response = result.Response
	resp.TotalResponse = result.TotalResponse
	resp.Search = []*pb.SearchDetail{}
	for _, detail := range result.Search {
		resp.Search = append(resp.Search, &pb.SearchDetail{
			Title:  detail.Title,
			Year:   detail.Year,
			ImdbID: detail.ImdbID,
			Type:   detail.Type,
			Poster: detail.Poster,
		})
	}

	return resp, nil
}

// GetMovie handles GRPC requests for get movie
func (srv *Service) GetMovie(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	resp := &pb.GetResponse{}
	// validate request
	if req == nil || req.ID == "" {
		log.Println("func GetMovie invalid request")
		return resp, errors.New("invalid request")
	}

	// get movie from usecase
	result, err := srv.movie.Get(entity.GetRequest{
		ID: req.ID,
	})
	if err != nil {
		log.Println("func GetMovie error from usecase", err)
		return resp, err
	}

	// parse result struct to proto
	resp.Title = result.Title
	resp.Year = result.Year
	resp.Rated = result.Rated
	resp.Released = result.Released
	resp.Runtime = result.Runtime
	resp.Genre = result.Genre
	resp.Director = result.Director
	resp.Writer = result.Writer
	resp.Actors = result.Actors
	resp.Plot = result.Plot
	resp.Language = result.Language
	resp.Country = result.Country
	resp.Awards = result.Awards
	resp.Poster = result.Poster
	resp.Ratings = []*pb.Rating{}
	for _, rating := range result.Ratings {
		resp.Ratings = append(resp.Ratings, &pb.Rating{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}
	resp.Metascore = result.Metascore
	resp.ImdbRating = result.ImdbRating
	resp.ImdbVotes = result.ImdbVotes
	resp.ImdbID = result.ImdbID
	resp.Type = result.Type
	resp.DVD = result.DVD
	resp.BoxOffice = result.BoxOffice
	resp.Production = result.Production
	resp.Website = result.Website
	resp.Response = result.Response

	return resp, nil
}
