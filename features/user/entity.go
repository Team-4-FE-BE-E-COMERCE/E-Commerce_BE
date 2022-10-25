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
	// Edit(input Core) (Core, error)
	GetDataId(param, token int) (data Core, err error)
	GetProfile(token int) (data Core, err error)
	PutDataId(data Core) (row int, err error)

	DeleteAkun(token int) (int, error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	// UpdateUser(input Core) (Core, error)
	// GetById(param, token int) (data UserCore, err error)
	SelectDataId(param, token int) (data Core, err error)
	MyProfile(token int) (data Core, err error)
	UpdateData(data Core) (row int, err error)
	DeleteData(token int) (int, error)
}
