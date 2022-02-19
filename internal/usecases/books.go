package usecases

import "github.com/Dimadetected/libraryBack/internal/models"

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
func (uc *UseCases) BooksTagsAdd(booksTags models.BookTags) error {
	return uc.r.BooksTagsAdd(booksTags)
}
func (uc *UseCases) BooksTagsDelete(booksTags models.BookTags) error {
	return uc.r.BooksTagsDelete(booksTags)
}
func (uc *UseCases) FavoriteBooksAdd(fb models.FavoriteBook) error {
	return uc.r.FavoriteBooksAdd(fb)
}
func (uc *UseCases) FavoriteBooksDelete(fb models.FavoriteBook) error {
	return uc.r.FavoriteBooksDelete(fb)
}
func (uc *UseCases) ProcessingBooksAdd(pb models.ProcessingBook) error {
	return uc.r.ProcessingBooksAdd(pb)
}
func (uc *UseCases) ProcessingBooksDelete(pb models.ProcessingBook) error {
	return uc.r.ProcessingBooksDelete(pb)
}

func (uc *UseCases) UpdateBook(id int, book models.Book) error {
	return uc.r.UpdateBook(id, book)
}
