package models

type Author struct {
	ID       int    `json:"id" db:"id" mapstructure:"id"`
	Name     string `json:"name" db:"name" mapstructure:"name"`
	Birthday string `json:"birthday" db:"birthday" mapstructure:"birthday"`
}
