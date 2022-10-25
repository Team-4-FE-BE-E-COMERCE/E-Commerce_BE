package service

import (
	"errors"
	// "log"
	"project/e-commerce/features/user"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostData(data user.Core) (int, error) {
	if data.Name == "" {
		return -1, errors.New("masukan nama ")
	}
	if data.Email == "" {
		return -1, errors.New("masukan email ")
	}
	if data.Password == "" {
		return -1, errors.New("masukann password")
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	data.Password = string(hashPass)
	row, err := usecase.userData.InsertData(data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *userUsecase) PutDataId(data user.Core) (int, error) {
	row, err := usecase.userData.UpdateData(data)
	return row, err
}

func (service *userUsecase) GetProfile(token int) (user.Core, error) {
	dataId, err := service.userData.MyProfile(token)
	if err != nil {
		return user.Core{}, err
	}
	return dataId, nil
}
func (service *userUsecase) GetDataId(param, token int) (user.Core, error) {

	dataId, err := service.userData.SelectDataId(param, token)
	if err != nil {
		return user.Core{}, err
	}

	return dataId, nil

}
func (usecase *userUsecase) DeleteAkun(token int) (int, error) {
	row, err := usecase.userData.DeleteData(token)
	if err != nil {
		return -1, err
	}
	return row, nil
}
