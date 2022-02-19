package usecases

import (
	"github.com/Dimadetected/libraryBack/internal/repositories"
)

type UseCases struct {
	r *repositories.Repositories
}

func NewUseCases(r *repositories.Repositories) *UseCases {
	return &UseCases{r: r}
}
