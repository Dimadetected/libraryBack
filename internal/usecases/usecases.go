package usecases

import (
	"github.com/Dimadetected/libraryBack/internal/models"
	"github.com/Dimadetected/libraryBack/internal/repositories"
)

type UseCases struct {
	r *repositories.Repositories
}

func NewUseCases(r *repositories.Repositories) *UseCases {
	return &UseCases{r: r}
}

func (uc *UseCases) GetBooks(limit, offset int) ([]models.Book, error) {
	return uc.r.GetBooks(limit, offset)
}
func (uc *UseCases) GetBook(id int) (models.Book, error) {
	return uc.r.GetBook(id)
}
func (uc *UseCases) DeleteBook(id int) error {
	return uc.r.DeleteBook(id)
}
func (uc *UseCases) StoreBook(book models.Book) (int64, error) {
	return uc.r.StoreBook(book)
}

func (uc *UseCases) UpdateBook(id int, book models.Book) error {
	return uc.r.UpdateBook(id, book)
}
