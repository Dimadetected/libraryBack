package repositories

import "github.com/Dimadetected/libraryBack/internal/models"

func (r *Repositories) GetUsers(limit, offset int) ([]models.User, error) {
	var Users []models.User
	if err := r.db.Select(&Users, `select * from users limit $1 offset $2`, limit, offset); err != nil {
		return nil, err
	}

	return Users, nil
}

func (r *Repositories) GetUser(id int) (models.User, error) {
	var User models.User
	if err := r.db.Get(&User, `select * from users where id = $1`, id); err != nil {
		return models.User{}, err
	}

	return User, nil
}
func (r *Repositories) DeleteUser(id int) error {
	if _, err := r.db.Exec(`delete from users where id = $1`, id); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) StoreUser(User models.User) (int64, error) {
	var id int64
	if err := r.db.QueryRow(`INSERT INTO users (name, email, birthday) 
									VALUES ($1,$2,$3) RETURNING id`, User.Name, User.Email, User.Birthday).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *Repositories) UpdateUser(id int, User models.User) error {
	if _, err := r.db.Exec(`UPDATE users 
									SET name = $1,
									    email = $2,
									    birthday = $3
									 	where id = $4`, User.Name, User.Email, User.Birthday, id); err != nil {
		return err
	}

	return nil
}
