package repositories

import (
	"fmt"
	"github.com/Dimadetected/libraryBack/internal/models"
)

func (r *Repositories) GetAuthors(limit, offset int) ([]models.Author, error) {
	var Authors []models.Author
	if err := r.db.Select(&Authors, `select * from Authors limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	fmt.Printf("2: %+v\n", Authors)
	return Authors, nil
}

func (r *Repositories) GetAuthor(id int) (models.Author, error) {
	var Author models.Author
	if err := r.db.Get(&Author, `select * from Authors where id = $1`, id); err != nil {
		return models.Author{}, err
	}

	return Author, nil
}
func (r *Repositories) DeleteAuthor(id int) error {
	if _, err := r.db.Exec(`delete from Authors where id = $1`, id); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) StoreAuthor(Author models.Author) (int64, error) {
	var id int64
	if err := r.db.QueryRow(`INSERT INTO authors (name, birthday) 
									VALUES ($1,$2) RETURNING id`, Author.Name, Author.Birthday).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *Repositories) UpdateAuthor(id int, Author models.Author) error {
	if _, err := r.db.Exec(`UPDATE Authors 
									SET name = $1,
									    birthday = $2
									 	where id = $3`, Author.Name, Author.Birthday, id); err != nil {
		return err
	}

	return nil
}
