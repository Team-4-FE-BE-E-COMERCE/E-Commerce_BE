package delivery

import "project/e-commerce/features/user"

type UserResponse struct {
	Name   string `json:"name" form:"name"`
	Images string `json:"images" form:"images"`
	Phone  int    `json:"phone" form:"phone"`
	Saldo  int    `json:"saldo"`
	Bio    string `json:"username" form:"username"`
	Email  string `json:"email" form:"email"`
}

func FromCore(data user.Core) UserResponse {
	return UserResponse{
		Name:   data.Name,
		Images: data.Images,
		Phone:  data.Phone,
		Bio:    data.Bio,
		Saldo:  int(data.Saldo),
		Email:  data.Email,
	}
}
