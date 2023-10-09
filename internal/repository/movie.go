package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"my-projects/stockbit/internal/constants/omdb"
	"my-projects/stockbit/internal/entity"
	"net/http"
)

// NewMovie to create new movie repository
func NewMovie(httpClient *http.Client) Movie {
	return &movieRP{
		httpClient,
	}
}

func (rp *movieRP) Search(request entity.SearchRequest) (entity.SearchResponse, error) {
	var resp entity.SearchResponse

	// construct HTTP request
	queryParams := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", omdb.URL, omdb.APIKey, request.SearchWord, request.Pagination)
	req, err := http.NewRequest(http.MethodGet, queryParams, nil)
	if err != nil {
		log.Println("func Search error create HTTP request", err)
		return resp, err
	}

	// do HTTP request
	httpResp, err := rp.httpClient.Do(req)
	if err != nil {
		log.Println("func Search error Do HTTP request", err)
		return resp, err
	}
	defer httpResp.Body.Close()

	// read all HTTP response
	raw, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Println("func Search error read HTTP response", err)
		return resp, err
	}
	fmt.Println("func Search success search from OMDB", string(raw))

	// parse HTTP response to struct
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		log.Println("func Search error parse HTTP response", err)
		return resp, err
	}

	// check if there is error from OMDB
	if resp.Response == "False" {
		log.Println("func Search error from OMDB", resp.Error)
		return resp, errors.New(resp.Error)
	}
	return resp, nil
}

func (rp *movieRP) Get(request entity.GetRequest) (entity.GetResponse, error) {
	var resp entity.GetResponse

	// construct HTTP request
	queryParams := fmt.Sprintf("%s/?apikey=%s&i=%s", omdb.URL, omdb.APIKey, request.ID)
	req, err := http.NewRequest(http.MethodGet, queryParams, nil)
	if err != nil {
		log.Println("func Get error create HTTP request", err)
		return resp, err
	}

	// do HTTP request
	httpResp, err := rp.httpClient.Do(req)
	if err != nil {
		log.Println("func Get error Do HTTP request", err)
		return resp, err
	}
	defer httpResp.Body.Close()

	// read all HTTP response
	raw, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Println("func Get error read HTTP response", err)
		return resp, err
	}
	fmt.Println("func Get success search from OMDB", string(raw))

	// parse HTTP response to struct
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		log.Println("func Get error parse HTTP response", err)
		return resp, err
	}

	// check if there is error from OMDB
	if resp.Response == "False" {
		log.Println("func Search error from OMDB", resp.Error)
		return resp, errors.New(resp.Error)
	}
	return resp, nil
}
