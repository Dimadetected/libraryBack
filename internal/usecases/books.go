package usecases

import (
	"fmt"
	"github.com/Dimadetected/libraryBack/internal/models"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"
)

func (uc *UseCases) GetBooks(limit, offset, authorID int, year, name string, tags string) ([]models.Book, error) {
	return uc.r.GetBooks(limit, offset, authorID, year, name, tags)
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
func (uc *UseCases) FavoriteBooksGet(id int) ([]models.FavoriteBook, error) {
	return uc.r.FavoriteBooksGet(id)
}
func (uc *UseCases) FavoriteBooksDelete(id int) error {
	return uc.r.FavoriteBooksDelete(id)
}
func (uc *UseCases) ProcessingBooksAdd(pb models.ProcessingBook) error {
	return uc.r.ProcessingBooksAdd(pb)
}
func (uc *UseCases) ProcessingBooksGet(userID int) ([]models.ProcessingBook, error) {
	return uc.r.ProcessingBooksGet(userID)
}
func (uc *UseCases) ProcessingBooksDelete(pb int) error {
	return uc.r.ProcessingBooksDelete(pb)
}

func (uc *UseCases) UpdateBook(id int, book models.Book) (int, error) {
	return uc.r.UpdateBook(id, book)
}
func (uc *UseCases) UpdateBookFile(id int, file multipart.File) error {
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Println(1)

	fileNameStart := strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(rand.Intn(1000000)) + ".jpg"
	fileName := "./photo/" + fileNameStart
	fileNameToWrite := "/photo/" + fileNameStart
	if err := ioutil.WriteFile(fileName, fileContents, 0664); err != nil {
		return err
	}
	fmt.Println(1)

	return uc.r.UpdateBookFile(id, fileNameToWrite)
}
