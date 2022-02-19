package handler

import (
	"github.com/Dimadetected/libraryBack/internal/models"
	"github.com/Dimadetected/libraryBack/pkg/respfmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) TagsGet(c *gin.Context) {
	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(c.Request.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	Tags, err := h.uc.GetTags(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Tags)
}
func (h *Handler) TagGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	Tag, err := h.uc.GetTag(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, Tag)
}
func (h *Handler) TagsDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err := h.uc.DeleteTag(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, id)
}
func (h *Handler) TagsUpdate(c *gin.Context) {
	var Tag models.Tag

	if err := c.BindJSON(&Tag); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err := h.uc.UpdateTag(id, Tag); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) TagsStore(c *gin.Context) {
	var Tag models.Tag

	if err := c.BindJSON(&Tag); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := h.uc.StoreTag(Tag)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"id": id,
	})
}
