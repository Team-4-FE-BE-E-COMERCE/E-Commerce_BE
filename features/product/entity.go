package product

import "time"

type User struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name_user" form:"name_user"`
}
type Core struct {
	ID        uint      `json:"id" form:"id"`
	Images    string    `json:"images" form:"images"`
	Name      string    `json:"name" form:"name"`
	Price     uint      `json:"price" form:"price"`
	Stock     uint      `json:"stock" form:"stock"`
	Detail    string    `json:"detail" form:"detail"`
	UserID    uint      `json:"user_id" form:"user_id"`
	User      User      `json:"user" form:"user"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type DataInterface interface {
	Insert(data Core) (Core, error)
	Edit(data Core, id uint) (Core, error)
	Remove(id uint) (Core, error)
	GetAll() ([]Core, error)
	GetByID(id uint) (Core, error)
	GetMy(token uint) ([]Core, error)
}

type UsecaseInterface interface {
	Create(data Core) (Core, error)
	Update(data Core, id uint) (Core, error)
	Delete(id uint) (Core, error)
	ShowAll() ([]Core, error)
	ShowByID(id uint) (Core, error)
	ShowMy(token uint) ([]Core, error)
}
