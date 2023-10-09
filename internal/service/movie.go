package service

import (
	"encoding/json"
	"log"
	logConst "my-projects/stockbit/internal/constants/log"
	"my-projects/stockbit/internal/entity"
	"my-projects/stockbit/internal/repository"
	"time"
)

// NewMovie to create new movie service
func NewMovie(movieRP repository.Movie, databaseRP repository.Database) Movie {
	return &movieSVC{
		movieRP,
		databaseRP,
	}
}

func (svc *movieSVC) Search(request entity.SearchRequest) (entity.SearchResponse, error) {
	// Search to third party
	resp, err := svc.movieRP.Search(request)
	if err != nil {
		log.Println("func Search error from repository", err)
		return resp, err
	}

	// Log search result and request
	go svc.log(request, resp, logConst.ActivitySearch)

	return resp, nil
}

func (svc *movieSVC) Get(request entity.GetRequest) (entity.GetResponse, error) {
	// Get from third party
	resp, err := svc.movieRP.Get(request)
	if err != nil {
		log.Println("func Get error from repository", err)
		return resp, err
	}

	// Log search result and request
	go svc.log(request, resp, logConst.ActivityGet)

	return resp, nil
}

func (svc *movieSVC) log(req interface{}, res interface{}, activity string) {
	rawReq, err := json.Marshal(req)
	if err != nil {
		log.Println("func log error parsing request", err)
		return
	}

	rawRes, err := json.Marshal(res)
	if err != nil {
		log.Println("func log error parsing response", err)
		return
	}

	err = svc.databaseRP.InsertLog(entity.Log{
		Activity:   activity,
		Request:    string(rawReq),
		Response:   string(rawRes),
		CreateTime: time.Now(),
	})
	if err != nil {
		log.Println("func log error insert", err)
	}
}
