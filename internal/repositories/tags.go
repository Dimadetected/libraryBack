package repositories

import (
	"github.com/Dimadetected/libraryBack/internal/models"
)

func (r *Repositories) GetTags(limit, offset int) ([]models.Tag, error) {
	var Tags []models.Tag
	if err := r.db.Select(&Tags, `select * from tags limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	return Tags, nil
}

func (r *Repositories) GetTag(id int) (models.Tag, error) {
	var Tag models.Tag
	if err := r.db.Get(&Tag, `select * from Tags where id = $1`, id); err != nil {
		return models.Tag{}, err
	}

	return Tag, nil
}
func (r *Repositories) DeleteTag(id int) error {
	if _, err := r.db.Exec(`delete from Tags where id = $1`, id); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) StoreTag(Tag models.Tag) (int64, error) {
	var id int64
	if err := r.db.QueryRow(`INSERT INTO tags (name) 
									VALUES ($1) RETURNING id`, Tag.Name).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *Repositories) UpdateTag(id int, Tag models.Tag) error {
	if _, err := r.db.Exec(`UPDATE Tags 
									SET name = $1
									 	where id = $2`, Tag.Name, id); err != nil {
		return err
	}

	return nil
}
