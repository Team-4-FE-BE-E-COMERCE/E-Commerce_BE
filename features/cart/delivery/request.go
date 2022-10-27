package delivery

import (
	"project/e-commerce/features/cart"
)

type CartRequest struct {
	Quantity  int `json:"quantity" form:"quantity"`
	ProductID int `json:"product_id" form:"product_id"`
	UserID    int `json:"user_id" form:"user_id"`
}

func toCore(data CartRequest) cart.Core {
	return cart.Core{
		Quantity:  data.Quantity,
		ProductID: uint(data.ProductID),
		UserID:    data.UserID,
	}
}
