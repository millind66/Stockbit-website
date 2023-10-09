package grpc

import (
	"context"
	"errors"
	"my-projects/stockbit/internal/entity"
	pb "my-projects/stockbit/proto"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
		name           string
		mock           mock
		request        *pb.SearchRequest
		expectedResult *pb.SearchResponse
		expectedError  error
	}{
		{
			name:           "request is nil",
			expectedError:  errors.New("invalid request"),
			expectedResult: &pb.SearchResponse{},
		},
		{
			name: "empty searchword",
			request: &pb.SearchRequest{
				Pagination: 1,
			},
			expectedError:  errors.New("invalid request"),
			expectedResult: &pb.SearchResponse{},
		},
		{
			name: "invalid pagination",
			request: &pb.SearchRequest{
				SearchWord: "batman",
			},
			expectedError:  errors.New("invalid request"),
			expectedResult: &pb.SearchResponse{},
		},
		{
			name: "error from usecase",
			request: &pb.SearchRequest{
				SearchWord: "batman",
				Pagination: 1,
			},
			mock: mock{
				Search: Search{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedError:  errors.New("foo"),
			expectedResult: &pb.SearchResponse{},
		},
		{
			name: "success",
			request: &pb.SearchRequest{
				SearchWord: "batman",
				Pagination: 1,
			},
			mock: mock{
				Search: Search{
					call: true,
					data: entity.SearchResponse{
						Search: []entity.SearchDetail{
							{
								ImdbID: "1",
							},
						},
					},
				},
			},
			expectedResult: &pb.SearchResponse{
				Search: []*pb.SearchDetail{
					&pb.SearchDetail{
						ImdbID: "1",
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			movieUC := NewMockMovie(mockCtrl)

			svc := NewGRPCService(movieUC)

			if tc.mock.Search.call {
				movieUC.EXPECT().Search(gomock.Any()).Return(tc.mock.Search.data, tc.mock.Search.err)
			}

			res, err := svc.SearchMovie(context.Background(), tc.request)
			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
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
		name           string
		mock           mock
		request        *pb.GetRequest
		expectedResult *pb.GetResponse
		expectedError  error
	}{
		{
			name:           "request is nil",
			expectedError:  errors.New("invalid request"),
			expectedResult: &pb.GetResponse{},
		},
		{
			name:           "empty ID",
			expectedError:  errors.New("invalid request"),
			expectedResult: &pb.GetResponse{},
		},
		{
			name: "error from usecase",
			request: &pb.GetRequest{
				ID: "1",
			},
			mock: mock{
				Get: Get{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedError:  errors.New("foo"),
			expectedResult: &pb.GetResponse{},
		},
		{
			name: "success",
			request: &pb.GetRequest{
				ID: "1",
			},
			mock: mock{
				Get: Get{
					call: true,
					data: entity.GetResponse{
						Ratings: []entity.Rating{
							{
								Source: "Rotten Tomatoes",
							},
						},
					},
				},
			},
			expectedResult: &pb.GetResponse{
				Ratings: []*pb.Rating{
					&pb.Rating{
						Source: "Rotten Tomatoes",
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			movieUC := NewMockMovie(mockCtrl)

			svc := NewGRPCService(movieUC)

			if tc.mock.Get.call {
				movieUC.EXPECT().Get(gomock.Any()).Return(tc.mock.Get.data, tc.mock.Get.err)
			}

			res, err := svc.GetMovie(context.Background(), tc.request)
			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
