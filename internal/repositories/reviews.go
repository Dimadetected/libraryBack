package repositories

import (
	"github.com/Dimadetected/libraryBack/internal/models"
)

func (r *Repositories) BooksReviewsGet(BookID, limit, offset int) ([]models.Review, error) {

	var reviews []models.Review
	if err := r.db.Select(&reviews, `select * from reviews where book_id = $1 order by id desc  limit $2 offset $3`, BookID, limit, offset); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *Repositories) BooksReviewGet(id int) (models.Review, error) {
	var review models.Review
	if err := r.db.Get(&review, `select * from Reviews where id = $1`, id); err != nil {
		return models.Review{}, err
	}

	return review, nil
}
func (r *Repositories) DeleteBooksReview(id int) error {
	if _, err := r.db.Exec(`delete from Reviews where id = $1`, id); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) BooksReviewStore(Review models.Review) (int64, error) {
	var id int64
	if err := r.db.QueryRow(`INSERT INTO Reviews (user_id, book_id, grade, positive, negative, description) 
									VALUES ($1,$2,$3, $4,$5, $6) RETURNING id`, Review.UserID, Review.BookID, Review.Grade, Review.Positive, Review.Negative, Review.Description).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *Repositories) UpdateReview(id int, Review models.Review) (int, error) {
	if _, err := r.db.Exec(`UPDATE Reviews 
									SET user_id = $1,
									    book_id = $2,
									    grade = $3, 
									    positive = $4,
									    negative = $5,
									    description = $6
									 	where id = $7`, Review.UserID, Review.BookID, Review.Grade, Review.Positive, Review.Negative, Review.Description, id); err != nil {
		return id, err
	}

	return id, nil
}
func (r *Repositories) BooksReviewsGradesAdd(Review models.ReviewGrades) error {
	var reviewFromDB models.ReviewGrades
	if err := r.db.Get(&reviewFromDB, `SELECT * from reviews_grades where review_id = $1`, Review.ReviewId); err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}

	if reviewFromDB.Id != 0 {
		if reviewFromDB.Status == 1 {
			if _, err := r.db.Exec(`UPDATE reviews r SET positive = positive - 1 where id = $1`, Review.ReviewId); err != nil {
				return err
			}
		} else if reviewFromDB.Status == 2 {
			if _, err := r.db.Exec(`UPDATE reviews r SET negative = negative - 1 where id = $1`, Review.ReviewId); err != nil {
				return err
			}
		}
		if _, err := r.db.Exec(`DELETE FROM reviews_grades where id = $1`, reviewFromDB.Id); err != nil {
			return err
		}
	}

	if Review.Status == 1 {
		if _, err := r.db.Exec(`UPDATE reviews r SET positive = positive + 1 where id = $1`, Review.ReviewId); err != nil {
			return err
		}
	} else {
		if _, err := r.db.Exec(`UPDATE reviews r SET negative = negative + 1 where id = $1`, Review.ReviewId); err != nil {
			return err
		}
	}

	if _, err := r.db.Exec(`INSERT INTO reviews_grades(review_id,user_id, book_id,status) VALUES ($1,$2,$3,$4) returning id`, Review.ReviewId, Review.UserID, Review.BookID, Review.Status); err != nil {
		return err
	}

	return nil
}
func (r *Repositories) BooksReviewsGradesGet(userID, bookID int) ([]models.ReviewGrades, error) {
	var grades []models.ReviewGrades
	if err := r.db.Select(&grades, `SELECT * from reviews_grades where book_id = $1 and user_id = $2`, bookID, userID); err != nil {
		return nil, err
	}

	return grades, nil
}
func (r *Repositories) BooksReviewsGradesDelete(id int) error {
	if _, err := r.db.Exec(`DELETE FROM reviews_grades where id = $1`, id); err != nil {
		return err
	}

	return nil
}
