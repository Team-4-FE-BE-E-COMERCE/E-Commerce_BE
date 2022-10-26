package delivery

import (
	"project/e-commerce/features/product"
	"time"
)

type ProductResponse struct {
	ID        uint         `json:"id" form:"id"`
	Images    string       `json:"images" form:"images"`
	Name      string       `json:"name" form:"name"`
	Price     uint         `json:"price" form:"price"`
	Stock     uint         `json:"stock" form:"stock"`
	Detail    string       `json:"detail" form:"detail"`
	UserID    uint         `json:"user_id" form:"user_id"`
	User      product.User `json:"user" form:"user"`
	CreatedAt time.Time    `json:"created_at" form:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" form:"updated_at"`
}

func toResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "product":
		cnv := core.(product.Core)
		res = ProductResponse{ID: cnv.ID, Images: cnv.Images, Name: cnv.Name, Price: cnv.Price, Stock: cnv.Stock, Detail: cnv.Detail, User: product.User{}, CreatedAt: cnv.CreatedAt, UpdatedAt: cnv.UpdatedAt}
	case "all":
		var res []ProductResponse
		cnv := core.([]product.Core)
		for _, val := range cnv {
			res = append(res, ProductResponse{ID: val.ID, Images: val.Images, Name: val.Name, Price: val.Price, Stock: val.Stock, Detail: val.Detail, User: product.User{}, CreatedAt: val.CreatedAt, UpdatedAt: val.UpdatedAt})
		}
	case "delete":
		cnv := core.(product.Core)
		res = ProductResponse{ID: cnv.ID}
	}
	return res
}

func successResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func failResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
