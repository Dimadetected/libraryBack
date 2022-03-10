package models

type Review struct {
	Id          int    `json:"id" db:"id"`
	UserID      int    `json:"user_id" db:"user_id"`
	Description string `json:"description" db:"description"`
	BookID      int    `json:"book_id" db:"book_id"`
	Grade       int    `json:"grade" db:"grade"`
	Positive    int    `json:"positive" db:"positive"`
	Negative    int    `json:"negative" db:"negative"`
}
type ReviewGrades struct {
	Id       int `json:"id" db:"id"`
	UserID   int `json:"user_id" db:"user_id"`
	BookID   int `json:"book_id" db:"book_id"`
	Status   int `json:"status" db:"status"`
	ReviewId int `json:"review_id" db:"review_id"`
}
