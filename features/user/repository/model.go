package repository

import (
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Images   string
	Phone    int
	Email    string `gorm:"unique"`
	Password string
	Bio      string
	Saldo    uint
}

func FromDomain(du user.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Saldo:    du.Saldo,
		Email:    du.Email,
		Password: du.Password,
		Images:   du.Images,
		Phone:    du.Phone,
		Bio:      du.Bio,
	}
}

func fromCore(dataCore user.Core) User {
	dataModel := User{
		Name:   dataCore.Name,
		Images: dataCore.Images,
		Bio:    dataCore.Bio,
		Phone:  dataCore.Phone,
		Saldo:  dataCore.Saldo,
		// Images:   dataCore.Images,
		// Phone:    dataCore.Phone,
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
		Bio:      data.Bio,
		Saldo:    data.Saldo,
	}

}
func ToDomain(u User) user.Core {
	return user.Core{
		ID:       u.ID,
		Name:     u.Name,
		Phone:    u.Phone,
		Password: u.Password,
		Images:   u.Images,
		Saldo:    u.Saldo,
		Bio:      u.Bio,
	}
}
func ToCoreList(data []User) []user.Core {
	var dataCore []user.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
