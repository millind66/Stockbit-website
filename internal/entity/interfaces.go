package entity

import (
	"database/sql"
)

type DBClient interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
