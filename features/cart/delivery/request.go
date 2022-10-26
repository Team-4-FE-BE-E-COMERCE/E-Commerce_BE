package delivery

import (
	"project/e-commerce/features/cart"
)

type CartRequest struct {
	Quantity  int
	ProductID int `json:"product_id" form:"product_id"`
	UserID    int
}

func toCore(data CartRequest) cart.Core {
	return cart.Core{
		Quantity:  data.Quantity,
		ProductID: uint(data.ProductID),
		UserID:    data.UserID,
	}
}
