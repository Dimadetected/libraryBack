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
	router.Static("/photo", "./photo")
	router.GET("/books", h.BooksGet)
	router.GET("/books/:id", h.BookGet)
	router.POST("/books/delete/:id", h.BooksDelete)
	router.POST("/books/store", h.BooksStore)
	router.POST("/books/update/:id", h.BooksUpdate)
	router.POST("/books/files/:id", h.BooksFilesUpdate)

	router.GET("/authors/:id", h.AuthorsGet)
	router.GET("/authors", h.AuthorsGet)
	router.POST("/authors/delete/:id", h.AuthorsDelete)
	router.POST("/authors/store", h.AuthorsStore)
	router.POST("/authors/update/:id", h.AuthorsUpdate)

	router.POST("books/tags", h.BooksTagsAdd)
	router.POST("books/tags/delete", h.BooksTagsDelete)

	router.GET("/tags/:id", h.TagGet)
	router.GET("/tags", h.TagsGet)
	router.POST("/tags/delete/:id", h.TagsDelete)
	router.POST("/tags/store", h.TagsStore)
	router.POST("/tags/update/:id", h.TagsUpdate)

	router.GET("books/favorite/:user_id", h.FavoriteBooksGet)
	router.POST("books/favorite", h.FavoriteBooksAdd)
	router.POST("books/favorite/delete/:id", h.FavoriteBooksDelete)

	router.GET("books/processing/:user_id", h.ProcessingBooksGet)
	router.POST("books/processing", h.ProcessingBooksAdd)
	router.POST("books/processing/delete/:id", h.ProcessingBooksDelete)

	router.GET("books/reviews", h.BooksReviewsGet)
	router.GET("books/review/:id", h.BooksReviewGet)
	router.POST("books/reviews/store", h.BooksReviewsAdd)
	router.POST("books/reviews/update/:id", h.BooksReviewsUpdate)
	router.POST("books/reviews/delete/:id", h.BooksReviewsDelete)

	router.GET("books/reviews/grades/:user_id/:book_id", h.BooksReviewsGradesGet)
	router.POST("books/reviews/grades", h.BooksReviewsGradesAdd)

	//users := router.Group("/users")
	//{
	//	users.GET("/:id", h.UsersGet)
	//	users.GET("/", h.UsersGet)
	//	users.DELETE("delete/:id", h.UsersDelete)
	//	users.POST("/store", h.UsersStore)
	//	users.PUT("/update/:id", h.UsersUpdate)
	//}

	if err := router.Run("localhost:8000"); err != nil {
		log.Fatal(err)
	}

}
