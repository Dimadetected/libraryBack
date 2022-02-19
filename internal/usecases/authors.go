package usecases

import "github.com/Dimadetected/libraryBack/internal/models"

func (uc *UseCases) GetAuthors(limit, offset int) ([]models.Author, error) {
	return uc.r.GetAuthors(limit, offset)
}
func (uc *UseCases) GetAuthor(id int) (models.Author, error) {
	return uc.r.GetAuthor(id)
}
func (uc *UseCases) DeleteAuthor(id int) error {
	return uc.r.DeleteAuthor(id)
}
func (uc *UseCases) StoreAuthor(Author models.Author) (int64, error) {
	return uc.r.StoreAuthor(Author)
}

func (uc *UseCases) UpdateAuthor(id int, Author models.Author) error {
	return uc.r.UpdateAuthor(id, Author)
}
