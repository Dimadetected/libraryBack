package usecases

import "github.com/Dimadetected/libraryBack/internal/models"

func (uc *UseCases) GetTags(limit, offset int) ([]models.Tag, error) {
	return uc.r.GetTags(limit, offset)
}
func (uc *UseCases) GetTag(id int) (models.Tag, error) {
	return uc.r.GetTag(id)
}
func (uc *UseCases) DeleteTag(id int) error {
	return uc.r.DeleteTag(id)
}
func (uc *UseCases) StoreTag(Tag models.Tag) (int64, error) {
	return uc.r.StoreTag(Tag)
}

func (uc *UseCases) UpdateTag(id int, Tag models.Tag) error {
	return uc.r.UpdateTag(id, Tag)
}
