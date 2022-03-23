package models

type User struct {
	ID       int    `json:"id" db:"id" mapstructure:"id"`
	Name     string `json:"name" db:"name" mapstructure:"name"`
	Email    string `json:"email" db:"email" mapstructure:"email"`
	Birthday string `json:"birthday" db:"birthday" mapstructure:"birthday"`
}

type UserRegister struct {
	Email    string `json:"email" db:"email" mapstructure:"email"`
	Password string `json:"password" db:"password" mapstructure:"password"`
}
