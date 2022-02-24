package models

type Book struct {
	ID          int    `json:"id" db:"id" mapstructure:"id"`
	Name        string `json:"name" db:"name" mapstructure:"name"`
	Description string `json:"description" db:"description" mapstructure:"description"`
	AuthorID    int    `json:"author_id" db:"author_id" mapstructure:"author_id"`
	Year        int    `json:"year" db:"year" mapstructure:"year"`
	Age         string `json:"age" db:"age" mapstructure:"age"`
}

type BookTags struct {
	BookID int `json:"book_id"`
	TagID  int `json:"tag_id"`
}
type FavoriteBook struct {
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
}
type ProcessingBook struct {
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
	Page   int `json:"page"`
}
