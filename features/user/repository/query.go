package repository

import (
	"errors"
	"log"
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) InsertData(data user.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *userData) UpdateData(data user.Core) (row int, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", data.ID).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return int(tx.RowsAffected), nil
}

func (repo *userData) product() []Product {

	var ProductMe []Product
	tx := repo.db.Find(&ProductMe) //.Preload("User")
	if tx.Error != nil {
		return nil
	}

	return ProductMe

}

func (repo *userData) MyProfile(token int) (user.Core, error) {

	var data User

	tx := repo.db.First(&data, token)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	allProduct := repo.product()
	dataPostUser := data.toUserCore(allProduct)

	log.Print(dataPostUser)

	return dataPostUser, nil

}

func (repo *userData) SelectDataId(param, token int) (user.Core, error) {

	var datacheck User
	txcheck := repo.db.Where("ID=?", token).First(&datacheck)
	if txcheck.Error != nil {
		return user.Core{}, errors.New("error tx")
	}

	if int(datacheck.ID) != token {
		return user.Core{}, errors.New("not have access")
	}

	var data User
	tx := repo.db.First(&data, param)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	allProduct := repo.product()
	dataPostUser := data.toUserCore(allProduct)

	return dataPostUser, nil

}
func (repo *userData) DeleteData(token int) (int, error) {

	tx := repo.db.Unscoped().Where("id = ?", token).Delete(&User{})
	if tx.Error != nil {
		return -1, tx.Error
	}

	return int(tx.RowsAffected), nil

}
