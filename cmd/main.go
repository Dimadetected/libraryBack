package main

import (
	"github.com/Dimadetected/libraryBack/internal/handler"
	"github.com/Dimadetected/libraryBack/internal/repositories"
	"github.com/Dimadetected/libraryBack/internal/usecases"
	"github.com/Dimadetected/libraryBack/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db, err := database.Connect("postgres", "postgres", "postgres", "localhost", "5432")
	if err != nil {
		log.Fatal(err)
	}

	r := repositories.NewRepositories(db)
	uc := usecases.NewUseCases(r)
	h := handler.NewHandler(uc)

	router := gin.Default()

	books := router.Group("/books")
	{
		books.GET("/:id", h.BookGet)
		books.GET("/", h.BooksGet)
		books.DELETE("delete/:id", h.BooksDelete)
		books.POST("/store", h.BooksStore)
		books.PUT("/update/:id", h.BooksUpdate)
	}

	if err := router.Run("localhost:8000"); err != nil {
		log.Fatal(err)
	}

}
