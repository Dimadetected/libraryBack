package models

type Tag struct {
	ID   int    `json:"id" db:"id" mapstructure:"id"`
	Name string `json:"name" db:"name" mapstructure:"name"`
}
