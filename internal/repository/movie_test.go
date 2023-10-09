package repository

import (
	"my-projects/stockbit/internal/entity"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tcs := []struct {
		name    string
		request entity.SearchRequest
		isError bool
	}{
		{
			name:    "failed",
			isError: true,
		},
		{
			name: "success",
			request: entity.SearchRequest{
				SearchWord: "batman",
				Pagination: 1,
			},
		},
	}

	httpClient := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			movieRP := NewMovie(httpClient)

			_, err := movieRP.Search(tc.request)
			if tc.isError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tcs := []struct {
		name    string
		request entity.GetRequest
		isError bool
	}{
		{
			name:    "failed",
			isError: true,
		},
		{
			name: "success",
			request: entity.GetRequest{
				ID: "tt0372784",
			},
		},
	}

	httpClient := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			movieRP := NewMovie(httpClient)

			_, err := movieRP.Get(tc.request)
			if tc.isError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
