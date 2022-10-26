package repository

import (
	"project/e-commerce/features/product"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.DataInterface {
	return &repoQuery{db: db}
}

func (rq *repoQuery) Insert(data product.Core) (product.Core, error) {
	var cnv Products = FromCore(data)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on insert product", err.Error())
		return product.Core{}, err
	}
	if err := rq.db.Preload("User").Find(&cnv).Error; err != nil {
		log.Error("error on join after insert", err.Error())
		return product.Core{}, err
	}
	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Edit(data product.Core, id uint) (product.Core, error) {
	var cnv Products = FromCore(data)
	if err := rq.db.Table("products").Where("id = ?", id).Updates(&cnv).Error; err != nil {
		log.Error("error on editing product", err.Error())
		return product.Core{}, err
	}
	if err := rq.db.Preload("User").Find(&cnv).Error; err != nil {
		log.Error("error on join after insert", err.Error())
		return product.Core{}, err
	}
	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Remove(id uint) (product.Core, error) {
	var data Products
	if err := rq.db.Delete(&data, "id = ?", id).Error; err != nil {
		log.Error("error on removing product", err.Error())
		return product.Core{}, err
	}
	res := ToCore(data)
	return res, nil
}

func (rq *repoQuery) GetAll() ([]product.Core, error) {
	var data []Products
	if err := rq.db.Preload("User").Find(&data).Error; err != nil {
		log.Error("error on get all product", err.Error())
		return nil, err
	}
	res := ToCoreArray(data)
	return res, nil
}

func (rq *repoQuery) GetByID(id uint) (product.Core, error) {
	var data Products
	if err := rq.db.Preload("User").First(&data, "id = ?", id).Error; err != nil {
		log.Error("error on get product by id", err.Error())
		return product.Core{}, err
	}
	res := ToCore(data)
	return res, nil
}

func (rq *repoQuery) GetMy(token uint) ([]product.Core, error) {
	var data []Products
	if err := rq.db.Preload("User").Find(&data, "user_id = ?", token).Error; err != nil {
		log.Error("error on getting my product", err.Error())
		return nil, err
	}
	res := ToCoreArray(data)
	return res, nil
}
