package order

import "time"

type Core struct {
	ID          uint
	Quantity    int
	TotalPrice  int
	OrderStatus string
	CartID      uint
	CreatedAt   time.Time
	Updated     time.Time
	DeletedAt   time.Time
}

type AddressCore struct {
	TransactionID uint
	Street        string
	City          string
	Province      string
	PostCode      uint
}

type PaymentCore struct {
	TransactionID uint
	Payment interface{}
}

type HistoryOrder struct {
	Images      string
	Name        string
	Price       uint
	Quantity    uint
	OrderStatus string
}

type OrderStatusCore struct {
	OrderStatus string
}

type UsecaseInterface interface {
	PostData(token int, data AddressCore, dataPay PaymentCore) (int, error)
	PutStatus(token int, status string) (int, error)
	DeleteOrder(token int, status string) (int, error)
	GetOrder(token int) ([]HistoryOrder, error)
}

type DataInterface interface {
	InsertData(token int, data AddressCore, dataPay PaymentCore) (int, error)
	UpdateStatus(token int, status string) (int, error)
	CancelOrder(token int, status string) (int, error)
	SelectOrder(token int) ([]HistoryOrder, error)
}
