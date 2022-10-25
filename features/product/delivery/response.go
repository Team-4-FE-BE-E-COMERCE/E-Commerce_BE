package delivery

import (
	"project/e-commerce/features/product"
	"time"
)

type ProductResponse struct {
	ID        uint
	Images    string
	Name      string
	Price     uint
	Qty       uint
	Detail    string
	UserID    uint
	User      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "product":
		cnv := core.(product.Core)
		res = ProductResponse{ID: cnv.ID, Images: cnv.Images, Name: cnv.Name, Price: cnv.Price, Qty: cnv.Qty, Detail: cnv.Detail, UserID: cnv.UserID, User: cnv.User.Name, CreatedAt: cnv.CreatedAt, UpdatedAt: cnv.UpdatedAt}
	case "all":
		var res []ProductResponse
		cnv := core.([]product.Core)
		for _, val := range cnv {
			res = append(res, ProductResponse{ID: val.ID, Images: val.Images, Name: val.Name, Price: val.Price, Qty: val.Qty, Detail: val.Detail, UserID: val.UserID, User: val.User.Name, CreatedAt: val.CreatedAt, UpdatedAt: val.UpdatedAt})
		}
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
