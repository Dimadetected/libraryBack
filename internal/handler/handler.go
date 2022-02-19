package handler

import (
	"github.com/Dimadetected/libraryBack/internal/usecases"
)

type Handler struct {
	uc *usecases.UseCases
}

func NewHandler(uc *usecases.UseCases) *Handler {
	return &Handler{
		uc: uc,
	}
}
