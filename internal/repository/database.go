package repository

import (
	"errors"
	"log"
	"my-projects/stockbit/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

// NewDatabase to create new database repository
func NewDatabase(dbClient entity.DBClient) Database {
	return &databaseRP{
		dbClient,
	}
}

func (rp *databaseRP) InsertLog(data entity.Log) error {
	// check if database client is nil
	if rp.dbClient == nil {
		log.Println("func InsertLog database client is nil")
		return errors.New("database client is nil")
	}
	query := "INSERT INTO activity_log(activity, request, response, create_time) VALUES (?, ?, ?, ?)"

	insert, err := rp.dbClient.Query(query, data.Activity, data.Request, data.Response, data.CreateTime)
	if err != nil {
		log.Println("func InsertLog error execute query", err)
		return err
	}
	insert.Close()

	return nil
}
