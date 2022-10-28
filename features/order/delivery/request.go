package delivery

import "project/e-commerce/features/order"

type Request struct {
	Street   string `json:"street" form:"street"`
	City     string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	PostCode uint   `json:"postcode" form:"postcode"`
}

func (data *Request) fromCore() order.AddressCore {
	dataAddress := order.AddressCore{
		Street:   data.Street,
		City:     data.City,
		Province: data.Province,
		PostCode: data.PostCode,
	}

	return dataAddress

}
