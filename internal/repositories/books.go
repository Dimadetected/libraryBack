package repositories

import (
	"fmt"
	"github.com/Dimadetected/libraryBack/internal/models"
	"strconv"
	"strings"
	"time"
)

func (r *Repositories) GetBooks(limit, offset, authorID int, year, name string, tags string) ([]models.Book, error) {

	var books []models.Book
	query := `select * from books`

	where := make([]string, 0)
	if authorID != 0 {
		where = append(where, "author_id = "+strconv.Itoa(authorID))
	}

	if year != "" {
		where = append(where, "year LIKE '%"+year+"%'")
	}

	if name != "" {
		where = append(where, "name LIKE '%"+name+"%'")
	}

	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}

	query += ` order by name limit $1 offset $2`

	if err := r.db.Select(&books, query, limit, offset); err != nil {
		return nil, err
	}
	var tagsArr []string
	if tags != "" {
		tagsArr = strings.Split(tags, ",")
	}

	var newBooks []models.Book

	for i := range books {
		var booksTags []models.BookTagsDB

		if err := r.db.Select(&booksTags, `select * from books_tags where book_id = $1`, books[i].ID); err != nil {
			return nil, err
		}

		countTags := 0
		for _, bt := range booksTags {
			for _, ta := range tagsArr {
				if strconv.Itoa(bt.TagID) == ta {
					countTags++
				}
			}
		}

		books[i].Tags = make([]int, 0)
		for _, bt := range booksTags {
			books[i].Tags = append(books[i].Tags, bt.TagID)
		}
		fmt.Println(countTags, len(tagsArr))
		if countTags == len(tagsArr) {
			newBooks = append(newBooks, books[i])
		}
	}

	return newBooks, nil
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
func (r *Repositories) FavoriteBooksGet(id int) ([]models.FavoriteBook, error) {
	favorites := make([]models.FavoriteBook, 0)
	if err := r.db.Select(&favorites, `SELECT * from favorite_books where user_id = $1`, id); err != nil {
		return nil, err
	}

	return favorites, nil
}
func (r *Repositories) FavoriteBooksDelete(id int) error {
	if _, err := r.db.Exec(`DELETE FROM favorite_books where id = $1`, id); err != nil {
		return err
	}

	return nil
}

func (r *Repositories) ProcessingBooksAdd(pb models.ProcessingBook) error {
	if _, err := r.db.Exec(`INSERT INTO processing_books (book_id,user_id, page, pages, created) VALUES($1,$2,$3, $4,$5)`, pb.BookID, pb.UserID, pb.Page, pb.Pages, time.Now().Format("2006-01-02 15:04")); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) ProcessingBooksGet(userID int) ([]models.ProcessingBook, error) {
	pb := make([]models.ProcessingBook, 0)
	if err := r.db.Select(&pb, `SELECT id,user_id,book_id,page,pages,to_char(created,'YYYY-MM-DD HH24:MI') as created from processing_books where user_id = $1 order by created desc`, userID); err != nil {
		return nil, err
	}

	return pb, nil
}
func (r *Repositories) ProcessingBooksDelete(pb int) error {
	if _, err := r.db.Exec(`DELETE FROM processing_books where id = $1`, pb); err != nil {
		return err
	}

	return nil
}
