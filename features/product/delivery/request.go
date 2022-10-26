package delivery

import (
	"project/e-commerce/features/product"
	"project/e-commerce/features/product/repository"
)

type ProductFormat struct {
	Images string          `json:"images" form:"images"`
	Name   string          `json:"name" form:"name"`
	Price  uint            `json:"price" form:"price"`
	Stock  uint            `json:"stock" form:"stock"`
	Detail string          `json:"detail" form:"detail"`
	UserID uint            `json:"user_id" form:"user_id"`
	User   repository.User `json:"user" form:"user" gorm:"foreignKey:UserID"`
}

func ToCore(i interface{}) product.Core {
	switch i.(type) {
	case ProductFormat:
		cnv := i.(ProductFormat)
		return product.Core{Images: cnv.Images, Name: cnv.Name, Price: cnv.Price, Stock: cnv.Stock, Detail: cnv.Detail, UserID: cnv.UserID}
	}
	return product.Core{}
}
