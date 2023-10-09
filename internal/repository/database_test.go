package repository

import (
	sql "database/sql"
	"errors"
	"my-projects/stockbit/internal/entity"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInsertLog(t *testing.T) {
	type Query struct {
		call bool
		rows *sql.Rows
		err  error
	}

	type mock struct {
		Query Query
	}

	tcs := []struct {
		name          string
		mock          mock
		request       entity.Log
		expectedError error
	}{
		{
			name: "error execute query",
			mock: mock{
				Query: Query{
					call: true,
					err:  errors.New("foo"),
				},
			},
			expectedError: errors.New("foo"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			dbClient := NewMockDBClient(mockCtrl)

			databaseRP := NewDatabase(dbClient)

			if tc.mock.Query.call {
				dbClient.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mock.Query.rows, tc.mock.Query.err)
			}

			err := databaseRP.InsertLog(tc.request)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
