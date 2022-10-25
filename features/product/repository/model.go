package repository

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Images   string
	Phone    string
	Email    string
	Password string
	Bio      string
	Saldo    uint
}

type Products struct {
	gorm.Model
	Image  string
	Name   string
	Price  uint
	Qty    uint
	Detail string
	UserID uint
	User   User
}

func FromCore(pc product.Core) Products {
	return Products{
		Model:  gorm.Model{ID: pc.ID, CreatedAt: pc.CreatedAt, UpdatedAt: pc.UpdatedAt},
		Image:  pc.Image,
		Name:   pc.Name,
		Price:  pc.Price,
		Qty:    pc.Qty,
		Detail: pc.Detail,
		UserID: pc.UserID,
		User:   User{Name: pc.Name},
	}
}

func ToCore(p Products) product.Core {
	return product.Core{
		ID:     p.ID,
		Image:  p.Image,
		Name:   p.Name,
		Price:  p.Price,
		Qty:    p.Qty,
		Detail: p.Detail,
		UserID: p.UserID,
		User:   product.User{Name: p.Name},
	}
}

func ToCoreArray(pa []Products) []product.Core {
	var res []product.Core
	for _, val := range pa {
		res = append(res, product.Core{ID: val.ID, Image: val.Image, Name: val.Name, Price: val.Price, Qty: val.Qty, Detail: val.Detail, UserID: val.UserID, User: product.User{Name: val.Name}})
	}
	return res
}
