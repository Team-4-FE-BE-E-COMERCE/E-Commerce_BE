package delivery

import "project/e-commerce/features/order"

type ResponHistoryOrder struct {
	Images      string `json:"images"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	TotalPrice  uint   `json:"total_price"`
	OrderStatus string `json:"order_status"`
}

func toRespon(data []order.HistoryOrder) []ResponHistoryOrder {
	var dataRes []ResponHistoryOrder
	for i, v := range data {
		dataRes = append(dataRes, ResponHistoryOrder{
			Images:      v.Images,
			Name:        v.Name,
			Price:       v.Price,
			Quantity:    v.Quantity,
			OrderStatus: v.OrderStatus,
		})

		dataRes[i].TotalPrice = v.Price * v.Quantity

	}

	return dataRes
}
