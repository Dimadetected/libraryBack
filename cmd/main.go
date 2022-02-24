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
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	})

	router.GET("/authors/:id", h.AuthorsGet)
	router.GET("/authors", h.AuthorsGet)
	router.POST("/authors/delete/:id", h.AuthorsDelete)
	router.POST("/authors/store", h.AuthorsStore)
	router.POST("/authors/update/:id", h.AuthorsUpdate)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	})
	//books := router.Group("/books")
	//{
	//	books.GET("/", h.BooksGet)
	//	books.GET("/:id", h.BookGet)
	//	books.DELETE("delete/:id", h.BooksDelete)
	//	books.POST("/store", h.BooksStore)
	//	books.PUT("/update/:id", h.BooksUpdate)
	//}
	//	tags := books.Group("tags")
	//	{
	//		tags.POST("/", h.BooksTagsAdd)
	//		tags.DELETE("/", h.BooksTagsDelete)
	//	}
	//
	//	favorite := books.Group("favorite")
	//	{
	//		favorite.POST("/", h.FavoriteBooksAdd)
	//		favorite.DELETE("/", h.FavoriteBooksDelete)
	//	}
	//
	//	processing := books.Group("processing")
	//	{
	//		processing.POST("/", h.ProcessingBooksAdd)
	//		processing.DELETE("/", h.ProcessingBooksDelete)
	//	}
	//}
	//
	//tags := router.Group("/tags")
	//{
	//	tags.GET("/:id", h.TagGet)
	//	tags.GET("/", h.TagsGet)
	//	tags.DELETE("delete/:id", h.TagsDelete)
	//	tags.POST("/store", h.TagsStore)
	//	tags.PUT("/update/:id", h.TagsUpdate)
	//}
	//authors := router.Group("/authors")
	//{
	//	authors.GET("/:id", h.AuthorsGet)
	//	authors.GET("/", h.AuthorsGet)
	//	authors.DELETE("delete/:id", h.AuthorsDelete)
	//	authors.POST("/store", h.AuthorsStore)
	//	authors.PUT("/update/:id", h.AuthorsUpdate)
	//}
	//
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
