package product

import "time"

type Core struct {
	ID        uint
	Image     string
	Name      string
	Price     uint
	Qty       uint
	Detail    string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Cores struct {
	ID        uint
	Image     string
	Name      string
	Price     uint
	Qty       uint
	Detail    string
	NameUser  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(data Core) (row uint, err error)
	Edit(data Core, id uint) (row uint, err error)
	Remove(id uint) (row uint, err error)
	GetAll() ([]Cores, error)
	GetByID(id uint) (Cores, error)
	GetMy(token uint) ([]Cores, error)
}

type UsecaseInterface interface {
	Create(data Core) (row uint, err error)
	Update(data Core, id uint) (row uint, err error)
	Delete(id uint) (row uint, err error)
	ShowAll() ([]Cores, error)
	ShowByID(id uint) (Cores, error)
	ShowMy(token uint) ([]Cores, error)
}
