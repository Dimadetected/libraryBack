package handler

import (
	"github.com/Dimadetected/libraryBack/internal/usecases"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc *usecases.UseCases
}

func NewHandler(uc *usecases.UseCases) *Handler {
	return &Handler{
		uc: uc,
	}
}

func setCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

}
