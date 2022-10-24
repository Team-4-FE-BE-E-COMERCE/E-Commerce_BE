package repository

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Image  string
	Name   string
	Price  uint
	Qty    uint
	Detail string
	UserID uint
}

type ProductsIt struct {
	gorm.Model
	Image    string
	Name     string
	Price    uint
	Qty      uint
	Detail   string
	NameUser string
}

func FromCore(pc product.Core) Products {
	return Products{
		Model:  gorm.Model{ID: pc.ID},
		Image:  pc.Image,
		Name:   pc.Name,
		Price:  pc.Price,
		Qty:    pc.Qty,
		Detail: pc.Detail,
		UserID: pc.UserID,
	}
}

func ToCore(p ProductsIt) product.Cores {
	return product.Cores{
		ID:       p.ID,
		Image:    p.Image,
		Name:     p.Name,
		Price:    p.Price,
		Qty:      p.Qty,
		Detail:   p.Detail,
		NameUser: p.NameUser,
	}
}

func ToDomainArray(pa []Products) []product.Core {
	var res []product.Core
	for _, val := range pa {
		res = append(res, product.Core{ID: val.ID, Image: val.Image, Name: val.Name, Price: val.Price, Qty: val.Qty, Detail: val.Detail, UserID: val.UserID})
	}
	return res
}

func ToDomainArrayIt(pi []ProductsIt) []product.Cores {
	var res []product.Cores
	for _, val := range pi {
		res = append(res, product.Cores{ID: val.ID, Image: val.Image, Name: val.Name, Price: val.Price, Qty: val.Qty, Detail: val.Detail, NameUser: val.NameUser})
	}
	return res
}
