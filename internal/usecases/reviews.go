package usecases

import (
	"github.com/Dimadetected/libraryBack/internal/models"
)

func (uc *UseCases) BooksReviewsGet(bookID, limit, offset int) ([]models.Review, error) {
	return uc.r.BooksReviewsGet(bookID, limit, offset)
}
func (uc *UseCases) BooksReviewGet(id int) (models.Review, error) {
	return uc.r.BooksReviewGet(id)
}
func (uc *UseCases) DeleteBookReview(id int) error {
	return uc.r.DeleteBooksReview(id)
}
func (uc *UseCases) StoreBookReview(review models.Review) (int64, error) {
	return uc.r.BooksReviewStore(review)
}
func (uc *UseCases) UpdateBookReview(id int, review models.Review) (int, error) {
	return uc.r.UpdateReview(id, review)
}
func (uc *UseCases) BooksReviewsGradesAdd(review models.ReviewGrades) error {
	return uc.r.BooksReviewsGradesAdd(review)
}
func (uc *UseCases) BooksReviewsGradesGet(userID, bookID int) ([]models.ReviewGrades, error) {
	return uc.r.BooksReviewsGradesGet(userID, bookID)
}
func (uc *UseCases) BooksReviewsGradesDelete(id int) error {
	return uc.r.BooksReviewsGradesDelete(id)
}
