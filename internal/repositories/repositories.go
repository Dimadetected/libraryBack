package repositories

import (
	"fmt"
	"github.com/Dimadetected/libraryBack/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repositories struct {
	db *sqlx.DB
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		db: db,
	}
}

func (r *Repositories) GetBooks(limit, offset int) ([]models.Book, error) {
	var books []models.Book
	if err := r.db.Select(&books, `select * from books limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	fmt.Printf("2: %+v\n", books)
	return books, nil
}

func (r *Repositories) GetBook(id int) (models.Book, error) {
	var book models.Book
	if err := r.db.Get(&book, `select * from books where id = $1`, id); err != nil {
		return models.Book{}, err
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
	if err := r.db.QueryRow(`INSERT INTO books (name, author_id, seria, year, page_count, format, type, weight, age) 
									VALUES ($1,$2,$3, $4,$5,$6,$7,$8,$9) RETURNING id`, book.Name, book.AuthorID, book.Seria, book.Year, book.PageCount, book.Format, book.Type, book.Weight, book.Age).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *Repositories) UpdateBook(id int, book models.Book) error {
	if _, err := r.db.Exec(`UPDATE books 
									SET name = $1,
									    author_id = $2,
									    seria = $3, 
									    year = $4,
									    page_count = $5,
									    format = $6,
									    type = $7,
									    weight = $8,
									    age = $9
									 	where id = $10`, book.Name, book.AuthorID, book.Seria, book.Year, book.PageCount, book.Format, book.Type, book.Weight, book.Age, id); err != nil {
		return err
	}

	return nil
}
