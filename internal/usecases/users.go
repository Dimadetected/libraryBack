package usecases

import "github.com/Dimadetected/libraryBack/internal/models"

func (uc *UseCases) GetUsers(limit, offset int) ([]models.User, error) {
	return uc.r.GetUsers(limit, offset)
}
func (uc *UseCases) UserRegister(reg *models.UserRegister) (int, error) {
	return uc.r.UserRegister(reg)
}
func (uc *UseCases) UserLogin(reg *models.UserRegister) (int, error) {
	return uc.r.UserLogin(reg)
}
func (uc *UseCases) GetUser(id int) (models.User, error) {
	return uc.r.GetUser(id)
}
func (uc *UseCases) DeleteUser(id int) error {
	return uc.r.DeleteUser(id)
}
func (uc *UseCases) StoreUser(User models.User) (int64, error) {
	return uc.r.StoreUser(User)
}

func (uc *UseCases) UpdateUser(id int, User models.User) error {
	return uc.r.UpdateUser(id, User)
}
