package api

import (
	"errors"
	"fmt"
	"my-projects/stockbit/internal/entity"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestSearchMovie(t *testing.T) {
	type Search struct {
		call bool
		data entity.SearchResponse
		err  error
	}

	type mock struct {
		Search Search
	}

	tcs := []struct {
		name       string
		mock       mock
		searchWord string
		pagination string
	}{
		{
			name:       "error parsing pagination",
			pagination: "asd",
			searchWord: "batman",
		},
		{
			name:       "invalid pagination",
			pagination: "0",
			searchWord: "batman",
		},
		{
			name:       "empty searchword",
			pagination: "1",
		},
		{
			name:       "error from usecase",
			pagination: "1",
			searchWord: "batman",
			mock: mock{
				Search: Search{
					call: true,
					err:  errors.New("foo"),
				},
			},
		},
		{
			name:       "success",
			pagination: "1",
			searchWord: "batman",
			mock: mock{
				Search: Search{
					call: true,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			movieUC := NewMockMovie(mockCtrl)

			svc := NewAPIService(movieUC)

			if tc.mock.Search.call {
				movieUC.EXPECT().Search(gomock.Any()).Return(tc.mock.Search.data, tc.mock.Search.err)
			}

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:9000/search?searchword=%s&pagination=%s", tc.searchWord, tc.pagination), nil)
			svc.SearchMovie(&dummyWriter{}, req)
		})
	}
}

func TestGetMovie(t *testing.T) {
	type Get struct {
		call bool
		data entity.GetResponse
		err  error
	}

	type mock struct {
		Get Get
	}

	tcs := []struct {
		name string
		mock mock
		id   string
	}{
		{
			name: "empty id",
		},
		{
			name: "error from usecase",
			id:   "1",
			mock: mock{
				Get: Get{
					call: true,
					err:  errors.New("foo"),
				},
			},
		},
		{
			name: "success",
			id:   "1",
			mock: mock{
				Get: Get{
					call: true,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			movieUC := NewMockMovie(mockCtrl)

			svc := NewAPIService(movieUC)

			if tc.mock.Get.call {
				movieUC.EXPECT().Get(gomock.Any()).Return(tc.mock.Get.data, tc.mock.Get.err)
			}

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:9000/get?id=%s", tc.id), nil)
			svc.GetMovie(&dummyWriter{}, req)
		})
	}
}
