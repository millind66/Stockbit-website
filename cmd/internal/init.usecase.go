package internal

import (
	"database/sql"
	"log"
	"my-projects/stockbit/internal/repository"
	"my-projects/stockbit/internal/service"
	"my-projects/stockbit/internal/usecase"
	"net/http"
	"time"
)

// GetUsecase to construct repository, service, and usecase
func GetUsecase() Usecases {
	// initialize single httpClient
	httpClient := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}

	// initialize database client
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stockbit")
	if err != nil {
		log.Println("func GetUsecase could not initialize database connection", err)
	}

	// initialize repository
	movieRP := repository.NewMovie(httpClient)
	databaseRP := repository.NewDatabase(db)

	// initialize service
	movieSVC := service.NewMovie(movieRP, databaseRP)

	// initialize usecase
	movieUC := usecase.NewMovie(movieSVC)
	return Usecases{
		Movie: movieUC,
	}
}
