package product

import "time"

type User struct {
	ID   uint `gorm:"primary_key"`
	Name string
}
type Core struct {
	ID        uint `gorm:"primary_key"`
	Images    string
	Name      string
	Price     uint
	Qty       uint
	Detail    string
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(data Core) (row uint, err error)
	Edit(data Core, id uint) (row uint, err error)
	Remove(id uint) (row uint, err error)
	GetAll() ([]Core, error)
	GetByID(id uint) (Core, error)
	GetMy(token uint) ([]Core, error)
}

type UsecaseInterface interface {
	Create(data Core) (row uint, err error)
	Update(data Core, id uint) (row uint, err error)
	Delete(id uint) (row uint, err error)
	ShowAll() ([]Core, error)
	ShowByID(id uint) (Core, error)
	ShowMy(token uint) ([]Core, error)
}
