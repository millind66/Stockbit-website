package service

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

	type InsertLog struct {
		call bool
		err  error
	}

	type mock struct {
		Search    Search
		InsertLog InsertLog
	}

	tcs := []struct {
		name           string
		mock           mock
		request        entity.SearchRequest
		expectedResult entity.SearchResponse
		expectedError  error
	}{
		{
			name: "error search from repo",
			mock: mock{
				Search: Search{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedError: errors.New("foo"),
		},
		{
			name: "success with error insert log",
			mock: mock{
				Search: Search{
					call: true,
					data: entity.SearchResponse{
						Response: "True",
					},
				},
				InsertLog: InsertLog{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedResult: entity.SearchResponse{
				Response: "True",
			},
		},
		{
			name: "success without error",
			mock: mock{
				Search: Search{
					call: true,
					data: entity.SearchResponse{
						Response: "True",
					},
				},
				InsertLog: InsertLog{
					call: true,
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

			movieRP := NewMockMovie(mockCtrl)
			databaseRP := NewMockDatabase(mockCtrl)

			movieSVC := NewMovie(movieRP, databaseRP)

			if tc.mock.Search.call {
				movieRP.EXPECT().Search(gomock.Any()).Return(tc.mock.Search.data, tc.mock.Search.err)
			}

			if tc.mock.InsertLog.call {
				databaseRP.EXPECT().InsertLog(gomock.Any()).Return(tc.mock.InsertLog.err).AnyTimes()
			}

			res, err := movieSVC.Search(tc.request)
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

	type InsertLog struct {
		call bool
		err  error
	}

	type mock struct {
		Get       Get
		InsertLog InsertLog
	}

	tcs := []struct {
		name           string
		mock           mock
		request        entity.GetRequest
		expectedResult entity.GetResponse
		expectedError  error
	}{
		{
			name: "error get from repo",
			mock: mock{
				Get: Get{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedError: errors.New("foo"),
		},
		{
			name: "success with error insert log",
			mock: mock{
				Get: Get{
					call: true,
					data: entity.GetResponse{
						Title: "Batman",
					},
				},
				InsertLog: InsertLog{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedResult: entity.GetResponse{
				Title: "Batman",
			},
		},
		{
			name: "success without error",
			mock: mock{
				Get: Get{
					call: true,
					data: entity.GetResponse{
						Title: "Batman",
					},
				},
				InsertLog: InsertLog{
					call: true,
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

			movieRP := NewMockMovie(mockCtrl)
			databaseRP := NewMockDatabase(mockCtrl)

			movieSVC := NewMovie(movieRP, databaseRP)

			if tc.mock.Get.call {
				movieRP.EXPECT().Get(gomock.Any()).Return(tc.mock.Get.data, tc.mock.Get.err)
			}

			if tc.mock.InsertLog.call {
				databaseRP.EXPECT().InsertLog(gomock.Any()).Return(tc.mock.InsertLog.err).AnyTimes()
			}

			res, err := movieSVC.Get(tc.request)
			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
