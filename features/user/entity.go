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
	Product  []ProductCore
}

type ProductCore struct {
	User_ID uint   `json:"user_id" form:"user_id"`
	ID      uint   `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Images  string `json:"images" form:"images"`
	Stock   int    `json:"stock" form:"stock"`
	Price   int    `json:"price" form:"price"`
	// User    Core
}

type UsecaseInterface interface {
	PostData(data Core) (row int, err error)
	GetDataId(param, token int) (data Core, err error)
	GetProfile(token int) (data Core, err error)
	PutDataId(data Core) (row int, err error)

	DeleteAkun(token int) (int, error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	SelectDataId(param, token int) (data Core, err error)
	MyProfile(token int) (data Core, err error)
	UpdateData(data Core) (row int, err error)
	DeleteData(token int) (int, error)
}
