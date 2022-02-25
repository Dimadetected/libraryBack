package repositories

import (
	"fmt"
	"github.com/Dimadetected/libraryBack/internal/models"
	"strings"
)

func (r *Repositories) GetBooks(limit, offset int) ([]models.Book, error) {

	var books []models.Book
	if err := r.db.Select(&books, `select * from books order by name limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	for i := range books {
		var booksTags []models.BookTagsDB
		if err := r.db.Select(&booksTags, `select * from books_tags where book_id = $1`, books[i].ID); err != nil {
			return nil, err
		}

		books[i].Tags = make([]int, 0)
		for _, bt := range booksTags {
			books[i].Tags = append(books[i].Tags, bt.TagID)
		}
	}

	return books, nil
}

func (r *Repositories) GetBook(id int) (models.Book, error) {
	var book models.Book
	if err := r.db.Get(&book, `select * from books where id = $1`, id); err != nil {
		return models.Book{}, err
	}

	var booksTags []models.BookTagsDB
	if err := r.db.Select(&booksTags, `select * from books_tags where book_id = $1`, book.ID); err != nil {
		return book, err
	}

	book.Tags = make([]int, 0)
	for _, bt := range booksTags {
		book.Tags = append(book.Tags, bt.TagID)
	}

	return book, nil
}
func (r *Repositories) DeleteBook(id int) error {
	if _, err := r.db.Exec(`delete from books where id = $1`, id); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) StoreBook(book models.Book) (int64, error) {
	var id int64
	if err := r.db.QueryRow(`INSERT INTO books (name, author_id, description, year, age) 
									VALUES ($1,$2,$3, $4,$5) RETURNING id`, book.Name, book.AuthorID, book.Description, book.Year, book.Age).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *Repositories) UpdateBook(id int, book models.Book) (int, error) {
	if _, err := r.db.Exec(`UPDATE books 
									SET name = $1,
									    author_id = $2,
									    description = $3, 
									    year = $4,
									    age = $5
									 	where id = $6`, book.Name, book.AuthorID, book.Description, book.Year, book.Age, id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *Repositories) BooksTagsAdd(booksTags models.BookTags) error {
	q := `INSERT INTO books_tags (book_id,tag_id) VALUES %s`
	var values []string
	for _, t := range booksTags.TagID {
		values = append(values, fmt.Sprintf("(%d, %d)", booksTags.BookID, t))
	}

	if _, err := r.db.Exec(fmt.Sprintf(q, strings.Join(values, ","))); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) UpdateBookFile(id int, fileName string) error {
	if _, err := r.db.Exec(`UPDATE books SET file = $1 WHERE id = $2`, fileName, id); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) BooksTagsDelete(booksTags models.BookTags) error {
	if _, err := r.db.Exec(`DELETE FROM books_tags where book_id = $1`, booksTags.BookID); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) FavoriteBooksAdd(fb models.FavoriteBook) error {
	if _, err := r.db.Exec(`INSERT INTO favorite_books (book_id,user_id) VALUES($1,$2)`, fb.BookID, fb.UserID); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) FavoriteBooksDelete(fb models.FavoriteBook) error {
	if _, err := r.db.Exec(`DELETE FROM favorite_books where book_id = $1 and user_id = $2`, fb.BookID, fb.UserID); err != nil {
		return err
	}

	return nil
}

func (r *Repositories) ProcessingBooksAdd(pb models.ProcessingBook) error {
	if _, err := r.db.Exec(`INSERT INTO processing_books (book_id,user_id, page) VALUES($1,$2,$3)`, pb.BookID, pb.UserID, pb.Page); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) ProcessingBooksDelete(pb models.ProcessingBook) error {
	if _, err := r.db.Exec(`DELETE FROM processing_books where book_id = $1 and user_id = $2`, pb.BookID, pb.UserID); err != nil {
		return err
	}

	return nil
}
