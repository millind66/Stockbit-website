package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"my-projects/stockbit/internal/entity"
	"my-projects/stockbit/internal/usecase"
	"net/http"
	"strconv"
)

// NewAPIService to create new API service
func NewAPIService(movieUC usecase.Movie) *Service {
	return &Service{
		movie: movieUC,
	}
}

// SearchMovie handles API requests for search movie
func (srv *Service) SearchMovie(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	searchword := queryValues.Get("searchword")
	paginationStr := queryValues.Get("pagination")

	w.Header().Set("Content-Type", "application/json")

	// parse pagination from string to integer
	pagination, err := strconv.Atoi(paginationStr)
	if err != nil {
		log.Println("func SearchMovie error parsing pagination", err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	// validate request
	if searchword == "" || pagination < 1 {
		log.Println("func SearchMovie invalid request")
		w.Write([]byte(fmt.Sprint(errors.New("invalid request"))))
		return
	}

	// search movie from usecase
	result, err := srv.movie.Search(entity.SearchRequest{
		SearchWord: searchword,
		Pagination: pagination,
	})
	if err != nil {
		log.Println("func SearchMovie error from usecase", err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	// parse result to JSON
	jsonbyte, err := json.Marshal(result)
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	w.Write(jsonbyte)
}

// GetMovie handles API requests for get movie
func (srv *Service) GetMovie(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	w.Header().Set("Content-Type", "application/json")

	// validate request
	if id == "" {
		log.Println("func GetMovie invalid request")
		w.Write([]byte(fmt.Sprint(errors.New("invalid request"))))
		return
	}

	// get movie from usecase
	result, err := srv.movie.Get(entity.GetRequest{
		ID: id,
	})
	if err != nil {
		log.Println("func GetMovie error from usecase", err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	// parse result to JSON
	jsonbyte, err := json.Marshal(result)
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	w.Write(jsonbyte)
}
