package usecase

import (
	"errors"
	entity "my-projects/stockbit/internal/entity"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	type Search struct {
		call bool
		data entity.SearchResponse
		err  error
	}

	type mock struct {
		Search Search
	}

	tcs := []struct {
		name           string
		mock           mock
		request        entity.SearchRequest
		expectedResult entity.SearchResponse
		expectedError  error
	}{
		{
			name: "error get from service",
			mock: mock{
				Search: Search{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedError: errors.New("foo"),
		},
		{
			name: "success",
			mock: mock{
				Search: Search{
					call: true,
					data: entity.SearchResponse{
						Response: "True",
					},
				},
			},
			expectedResult: entity.SearchResponse{
				Response: "True",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			movieSVC := NewMockMovie(mockCtrl)

			movieUC := NewMovie(movieSVC)

			if tc.mock.Search.call {
				movieSVC.EXPECT().Search(gomock.Any()).Return(tc.mock.Search.data, tc.mock.Search.err)
			}

			res, err := movieUC.Search(tc.request)
			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGet(t *testing.T) {
	type Get struct {
		call bool
		data entity.GetResponse
		err  error
	}

	type mock struct {
		Get Get
	}

	tcs := []struct {
		name           string
		mock           mock
		request        entity.GetRequest
		expectedResult entity.GetResponse
		expectedError  error
	}{
		{
			name: "error get from service",
			mock: mock{
				Get: Get{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedError: errors.New("foo"),
		},
		{
			name: "success",
			mock: mock{
				Get: Get{
					call: true,
					data: entity.GetResponse{
						Title: "Batman",
					},
				},
			},
			expectedResult: entity.GetResponse{
				Title: "Batman",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			movieSVC := NewMockMovie(mockCtrl)

			movieUC := NewMovie(movieSVC)

			if tc.mock.Get.call {
				movieSVC.EXPECT().Get(gomock.Any()).Return(tc.mock.Get.data, tc.mock.Get.err)
			}

			res, err := movieUC.Get(tc.request)
			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
