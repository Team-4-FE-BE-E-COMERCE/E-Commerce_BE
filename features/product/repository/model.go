package repository

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name_user" form:"name_user"`
}

type Products struct {
	gorm.Model
	Images string `json:"images" form:"images"`
	Name   string `json:"name" form:"name"`
	Price  uint   `json:"price" form:"price"`
	Stock  uint   `json:"stock" form:"stock"`
	Detail string `json:"detail" form:"detail"`
	UserID uint   `json:"user_id" form:"user_id"`
	User   User   `gorm:"foreignKey:UserID"`
}

func FromCore(pc product.Core) Products {
	return Products{
		Model:  gorm.Model{ID: pc.ID, CreatedAt: pc.CreatedAt, UpdatedAt: pc.UpdatedAt},
		Images: pc.Images,
		Name:   pc.Name,
		Price:  pc.Price,
		Stock:  pc.Stock,
		Detail: pc.Detail,
		UserID: pc.UserID,
	}
}

func ToCore(p Products) product.Core {
	return product.Core{
		ID:        p.ID,
		Images:    p.Images,
		Name:      p.Name,
		Price:     p.Price,
		Stock:     p.Stock,
		Detail:    p.Detail,
		UserID:    p.UserID,
		User:      product.User(p.User),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToCoreArray(pa []Products) []product.Core {
	var res []product.Core
	for _, val := range pa {
		res = append(res, product.Core{ID: val.ID, Images: val.Images, Name: val.Name, Price: val.Price, Stock: val.Stock, Detail: val.Detail, UserID: val.UserID, User: product.User(val.User), CreatedAt: val.CreatedAt, UpdatedAt: val.UpdatedAt})
	}
	return res
}
