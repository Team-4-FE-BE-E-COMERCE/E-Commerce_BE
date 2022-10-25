package delivery

import (
	"project/e-commerce/features/product"
)

type ProductFormat struct {
	Image  string `json:"image" form:"image"`
	Name   string `json:"name" form:"name"`
	Price  uint   `json:"price" form:"price"`
	Qty    uint   `json:"qty" form:"qty"`
	Detail string `json:"detail" form:"detail"`
	UserID uint   `json:"user" form:"user"`
}

func ToCore(i interface{}) product.Core {
	switch i.(type) {
	case ProductFormat:
		cnv := i.(ProductFormat)
		return product.Core{Image: cnv.Image, Name: cnv.Name, Price: cnv.Price, Qty: cnv.Qty, Detail: cnv.Detail, UserID: cnv.UserID}
	}
	return product.Core{}
}
