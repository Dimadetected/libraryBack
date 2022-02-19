package models

type Book struct {
	ID        int    `json:"id" db:"id" mapstructure:"id"`
	Name      string `json:"name" db:"name" mapstructure:"name"`
	AuthorID  int    `json:"author_id" db:"author_id" mapstructure:"author_id"`
	Seria     string `json:"seria" db:"seria" mapstructure:"seria"`
	Year      int    `json:"year" db:"year" mapstructure:"year"`
	PageCount int    `json:"page_count" db:"page_count" mapstructure:"page_count"`
	Format    string `json:"format" db:"format" mapstructure:"format"`
	Type      string `json:"type" db:"type" mapstructure:"type"`
	Weight    int    `json:"weight" db:"weight" mapstructure:"weight"`
	Age       string `json:"age" db:"age" mapstructure:"age"`
}
