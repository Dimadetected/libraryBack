package handler

import (
	"github.com/Dimadetected/libraryBack/internal/models"
	"github.com/Dimadetected/libraryBack/pkg/respfmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) AuthorsGet(c *gin.Context) {
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

	Authors, err := h.uc.GetAuthors(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Authors)
}
func (h *Handler) AuthorGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	Author, err := h.uc.GetAuthor(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Author)
}
func (h *Handler) AuthorsDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err := h.uc.DeleteAuthor(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, id)
}
func (h *Handler) AuthorsUpdate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var Author models.Author

	if err := c.BindJSON(&Author); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err := h.uc.UpdateAuthor(id, Author); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) AuthorsStore(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var Author models.Author

	if err := c.BindJSON(&Author); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := h.uc.StoreAuthor(Author)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"id": id,
	})
}
