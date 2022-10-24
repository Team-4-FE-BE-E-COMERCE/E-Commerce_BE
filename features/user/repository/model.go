package repository

import (
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Images   string
	Phone    string
	Email    string `gorm:"unique"`
	Password string
}

func fromCore(dataCore user.Core) User {
	dataModel := User{
		Name:     dataCore.Name,
		Images:   dataCore.Images,
		Phone:    dataCore.Phone,
		Email:    dataCore.Email,
		Password: dataCore.Password,
	}
	return dataModel

}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Images:   data.Images,
		Phone:    data.Phone,
		Email:    data.Email,
		Password: data.Password,
	}

}

func ToCoreList(data []User) []user.Core {
	var dataCore []user.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
