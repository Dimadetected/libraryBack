package handler

import (
	"github.com/Dimadetected/libraryBack/internal/models"
	"github.com/Dimadetected/libraryBack/pkg/respfmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) BooksGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(c.Request.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	books, err := h.uc.GetBooks(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, books)
}
func (h *Handler) BookGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	book, err := h.uc.GetBook(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, book)
}
func (h *Handler) BooksDelete(c *gin.Context) {
	setCors(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err := h.uc.DeleteBook(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, id)
}
func (h *Handler) BooksUpdate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err := h.uc.UpdateBook(id, book); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) BooksStore(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := h.uc.StoreBook(book)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) BooksTagsAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.BookTags

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.BooksTagsAdd(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) BooksTagsDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.BookTags

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.BooksTagsDelete(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) FavoriteBooksAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.FavoriteBook

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.FavoriteBooksAdd(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) FavoriteBooksDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.FavoriteBook

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.FavoriteBooksDelete(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) ProcessingBooksAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.ProcessingBook

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.ProcessingBooksAdd(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) ProcessingBooksDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.ProcessingBook

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.ProcessingBooksDelete(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
