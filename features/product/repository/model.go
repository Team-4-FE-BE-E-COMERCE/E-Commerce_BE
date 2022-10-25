package repository

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

type Products struct {
	gorm.Model
	Images string
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
		Images: pc.Images,
		Name:   pc.Name,
		Price:  pc.Price,
		Qty:    pc.Qty,
		Detail: pc.Detail,
		UserID: pc.UserID,
		User:   User{ID: pc.ID, Name: pc.Name},
	}
}

func ToCore(p Products) product.Core {
	return product.Core{
		ID:        p.ID,
		Images:    p.Images,
		Name:      p.Name,
		Price:     p.Price,
		Qty:       p.Qty,
		Detail:    p.Detail,
		UserID:    p.UserID,
		User:      product.User{ID: p.ID, Name: p.Name},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToCoreArray(pa []Products) []product.Core {
	var res []product.Core
	for _, val := range pa {
		res = append(res, product.Core{ID: val.ID, Images: val.Images, Name: val.Name, Price: val.Price, Qty: val.Qty, Detail: val.Detail, UserID: val.UserID, User: product.User{ID: val.ID, Name: val.Name}})
	}
	return res
}
