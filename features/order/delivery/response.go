package delivery

import "project/e-commerce/features/order"

type ResponHistoryOrder struct {
	Images      string
	Name        string
	Price       uint
	Quantity    uint
	TotalPrice  uint
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
