package repository

import (
	"project/e-commerce/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToDomain(u User) login.Core {
	return login.Core{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
	}
}

func FromDomain(du login.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Password: du.Password,
	}
}
