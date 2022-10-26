package repository

import (
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Images   string
// 	Phone    int
// 	Email    string `gorm:"unique"`
// 	Password string
// 	Bio      string
// 	Saldo    uint
// 	Product  []Product //`gorm:"foreignKey:ProductID"`
// }

// type Product struct {
// 	ID     uint   `json:"id" form:"id"`
// 	Images string `json:"images" form:"images"`
// 	Name   string `json:"name" form:"name"`
// 	Price  int    `json:"price" form:"price"`
// 	Stock  int    `json:"stock" form:"stock"`
// 	// User   User
// }

// func FromDomain(du user.Core) User {
// 	return User{
// 		Model:    gorm.Model{ID: du.ID},
// 		Name:     du.Name,
// 		Saldo:    du.Saldo,
// 		Email:    du.Email,
// 		Password: du.Password,
// 		Images:   du.Images,
// 		Phone:    du.Phone,
// 		Bio:      du.Bio,
// 	}
// }

// func fromCore(dataCore user.Core) User {
// 	dataModel := User{
// 		Name:   dataCore.Name,
// 		Images: dataCore.Images,
// 		Bio:    dataCore.Bio,
// 		Phone:  dataCore.Phone,
// 		Saldo:  dataCore.Saldo,
// 		// Images:   dataCore.Images,
// 		// Phone:    dataCore.Phone,
// 		Email:    dataCore.Email,
// 		Password: dataCore.Password,
// 	}
// 	return dataModel

// }

// func (data *User) toCore() user.Core {
// 	return user.Core{
// 		ID:       data.ID,
// 		Name:     data.Name,
// 		Images:   data.Images,
// 		Phone:    data.Phone,
// 		Email:    data.Email,
// 		Password: data.Password,
// 		Bio:      data.Bio,
// 		Saldo:    data.Saldo,
// 	}

// }
// func ToDomain(u User) user.Core {
// 	return user.Core{
// 		ID:       u.ID,
// 		Name:     u.Name,
// 		Phone:    u.Phone,
// 		Password: u.Password,
// 		Images:   u.Images,
// 		Saldo:    u.Saldo,
// 		Bio:      u.Bio,
// 	}
// }
// func ToCoreList(data []User) []user.Core {
// 	var dataCore []user.Core
// 	for key := range data {
// 		dataCore = append(dataCore, data[key].toCore())
// 	}
// 	return dataCore
// }

// func (data *User) toUserCore(products []Product) user.Core {

// 	dataUser := user.Core{
// 		ID:     data.ID,
// 		Name:   data.Name,
// 		Email:  data.Email,
// 		Images: data.Images,
// 		Phone:  data.Phone,

// 		// Saldo:  data.Saldo,

// 		// Created_At: data.CreatedAt,
// 		// Updated_At: data.UpdatedAt,
// 	}

// 	commentUser := toCoreListProd(products)

// 	for i, v := range commentUser {
// 		if v.User_ID == dataUser.ID {
// 			dataUser.Product = append(dataUser.Product, commentUser[i])
// 		}
// 	}

// 	return dataUser
// }

// func (data *Product) toCoreProd() user.ProductCore {
// 	return user.ProductCore{
// 		// User_ID: data.User_ID,
// 		Name:   data.Name,
// 		Images: data.Images,
// 		Price:  data.Price,
// 		Stock:  data.Stock,
// 		// Description: data.Description,
// 		// UserID:      int(data.UserID),
// 	}
// }

// func toCoreListProd(data []Product) []user.ProductCore {
// 	var dataCore []user.ProductCore
// 	for key := range data {
// 		dataCore = append(dataCore, data[key].toCoreProd())
// 	}
// 	return dataCore
// }

type User struct {
	gorm.Model
	Name     string
	Images   string
	Phone    int
	Email    string `gorm:"unique"`
	Password string
	Bio      string
	Saldo    uint
	Product  []Product
}

type Product struct {
	gorm.Model
	User_ID uint `json:"user_id"`
	// ID      uint   `json:"id"`
	Name   string `json:"name"`
	Images string `json:"images"`
	Price  int    `json:"price"`
	Stock  int    `json:"stock"`
	// Products []user.ProductCore
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
		// ID:     dataCore.ID,
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
func (data *Product) toCoreProd() user.ProductCore {
	return user.ProductCore{
		User_ID: data.User_ID,
		ID:      data.ID,
		Name:    data.Name,
		Images:  data.Images,
		Price:   data.Price,
		Stock:   data.Stock,
		// Description: data.Description,
		// UserID:      int(data.UserID),
	}
}

func toCoreListProd(data []Product) []user.ProductCore {
	var dataCore []user.ProductCore
	for key := range data {
		dataCore = append(dataCore, data[key].toCoreProd())
	}
	return dataCore
}

func (data *User) toUserCore(products []Product) user.Core {

	dataUser := user.Core{
		ID:     data.ID,
		Name:   data.Name,
		Email:  data.Email,
		Images: data.Images,
		Bio:    data.Bio,
		Phone:  data.Phone,
	}

	product := toCoreListProd(products)

	for i, v := range product {
		if v.User_ID == dataUser.ID {
			dataUser.Product = append(dataUser.Product, product[i])
		}
	}

	return dataUser
}

func ToCoreArray(pa []User) []user.Core {
	var res []user.Core
	for _, val := range pa {
		res = append(res, user.Core{ID: val.ID, Images: val.Images, Name: val.Name, Bio: val.Bio, Email: val.Email, Phone: val.Phone, Product: val.toCore().Product})
	}
	return res
}
