package service

import (
	"errors"
	"project/e-commerce/features/product"
	"strings"

	"github.com/labstack/gommon/log"
)

type productService struct {
	qry product.DataInterface
}

func New(repo product.DataInterface) product.UsecaseInterface {
	return &productService{qry: repo}
}

func (ps *productService) Create(data product.Core) (product.Core, error) {
	res, err := ps.qry.Insert(data)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return product.Core{}, errors.New("rejected from database")
		}
		return product.Core{}, errors.New("some problem on database")
	}
	return res, nil
}

func (ps *productService) Update(data product.Core, id uint) (product.Core, error) {
	res, err := ps.qry.Edit(data, id)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return product.Core{}, errors.New("rejected from database")
		}
		return product.Core{}, errors.New("some problem on database")
	}
	return res, nil
}

func (ps *productService) Delete(id uint) (product.Core, error) {
	res, err := ps.qry.Remove(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return product.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return product.Core{}, errors.New("no data")
		}
	}
	return res, nil
}

func (ps *productService) ShowAll() ([]product.Core, error) {
	res, err := ps.qry.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}
	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}
	return res, nil
}

func (ps *productService) ShowByID(id uint) (product.Core, error) {
	res, err := ps.qry.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return product.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return product.Core{}, errors.New("no data")
		}
	}
	return res, nil
}

func (ps *productService) ShowMy(token uint) ([]product.Core, error) {
	res, err := ps.qry.GetMy(token)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}
	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}
	return res, nil
}
