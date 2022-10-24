package user

type Core struct {
	ID       uint
	Name     string
	Images   string
	Phone    string
	Email    string
	Password string
	Bio      string
	Saldo    uint
}

type UsecaseInterface interface {
	PostData(data Core) (row int, err error)
	GetData(id int) (data Core, err error)
	GetProfile(token int) (data Core, err error)
	PutDataId(data Core, id int) (row int, err error)
	DeleteAkun(id int) (row int, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	SelectData(id int) (data Core, err error)
	MyProfile(token int) (data Core, err error)
	UpdateData(data Core, id int) (row int, err error)
	DeleteData(id int) (row int, err error)
}
