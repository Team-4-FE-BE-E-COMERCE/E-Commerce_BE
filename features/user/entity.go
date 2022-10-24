package user

type Core struct {
	ID       uint
	Name     string
	Images   string
	Phone    int
	Email    string
	Password string
	Bio      string
	Saldo    uint
}

type UsecaseInterface interface {
	PostData(data Core) (row int, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
}
