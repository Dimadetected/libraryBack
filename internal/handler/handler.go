package handler

import (
	"github.com/Dimadetected/libraryBack/internal/models"
	"github.com/Dimadetected/libraryBack/internal/usecases"
	"github.com/Dimadetected/libraryBack/pkg/respfmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	uc *usecases.UseCases
}

func NewHandler(uc *usecases.UseCases) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) BooksGet(c *gin.Context) {
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
