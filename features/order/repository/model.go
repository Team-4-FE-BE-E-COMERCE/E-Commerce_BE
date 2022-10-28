package data

import (
	"project/e-commerce/features/order"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Quantity    int
	TotalPrice  int
	OrderStatus string
	CartID      uint
}

type Payment struct {
	TransactionID uint `gorm:"primary_key;ForeignKey:TransactionID"`
	Payment       string
}

type Address struct {
	TransactionID uint `gorm:"primary_key;ForeignKey:TransactionID"`
	Street        string
	City          string
	Province      string
	PostCode      uint
}

type Product struct {
	gorm.Model
	Name         string
	Images       string
	Price        int
	Stock        int
	Desc         string
	UserID       uint
	CategoriesID uint
	CartID       []Cart `gorm:"foreignKey:ProductID"`
}

type Cart struct {
	gorm.Model
	Quantity  int
	ProductID uint
	UserID    uint
}

type Results struct {
	ID          uint
	Quantity    int
	Name        string
	Images      string
	Price       int
	OrderStatus string
	UserID      uint
	ProductID   uint
}

type result struct {
	CartID      uint
	Quantity    int
	TotalPrice  int
	OrderStatus string
}

type DBTransaction struct {
	Cart_id  uint
	Quantity int
	Stock    int
	ID       uint
}

func insertJoin(data []Results) []result {
	var dataRes []result
	for _, v := range data {
		dataRes = append(dataRes, result{
			Quantity:   v.Quantity,
			TotalPrice: v.Price,
			CartID:     v.ID,
			// OrderStatus: v.OrderStatus,
		})
	}
	return dataRes

}

func FromCore(data order.Core) Transaction {
	return Transaction{
		Model:       gorm.Model{ID: data.ID, CreatedAt: data.CreatedAt, UpdatedAt: data.Updated},
		Quantity:    data.Quantity,
		TotalPrice:  data.TotalPrice,
		OrderStatus: data.OrderStatus,
		CartID:      data.CartID,
	}
}

func insert(res []result) []Transaction {
	var data []Transaction
	for _, v := range res {
		data = append(data, Transaction{
			Quantity:    v.Quantity,
			TotalPrice:  v.TotalPrice,
			CartID:      v.CartID,
			OrderStatus: "waiting",
			// OrderStatus: v.OrderStatus,
		})
	}
	return data

}

func (tx *Transaction) toCore() order.Core {
	return order.Core{
		ID:          tx.ID,
		Quantity:    tx.Quantity,
		TotalPrice:  tx.TotalPrice,
		OrderStatus: tx.OrderStatus,
		CartID:      tx.CartID,
		CreatedAt:   tx.CreatedAt,
		Updated:     tx.UpdatedAt,
		DeletedAt:   tx.DeletedAt.Time,
	}
}

func toDb(data order.AddressCore) Address {
	return Address{
		TransactionID: data.TransactionID,
		Street:        data.Street,
		City:          data.City,
		Province:      data.Province,
		PostCode:      data.PostCode,
	}

}

func toDbPay(data order.PaymentCore) Payment {
	return Payment{
		Payment: data.Payment,
	}

}
