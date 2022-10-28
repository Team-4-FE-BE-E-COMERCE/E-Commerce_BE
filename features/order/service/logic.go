package service

import (
	"errors"
	"project/e-commerce/features/order"
)

type transactionUsecase struct {
	transactionData order.DataInterface
}

func New(data order.DataInterface) order.UsecaseInterface {
	return &transactionUsecase{
		transactionData: data,
	}
}

func (usecase *transactionUsecase) PostData(token int, data order.AddressCore, dataPay order.PaymentCore) (int, error) {

	if data.City == "" || data.PostCode == 0 || data.Province == "" || data.Street == "" {
		return -1, errors.New("error")
	}

	row, err := usecase.transactionData.InsertData(token, data, dataPay)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *transactionUsecase) PutStatus(token int, status string) (int, error) {

	row, err := usecase.transactionData.UpdateStatus(token, status)
	if err != nil {
		return -1, err
	}

	return row, nil

}

func (usecase *transactionUsecase) DeleteOrder(token int, status string) (int, error) {
	row, err := usecase.transactionData.CancelOrder(token, status)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *transactionUsecase) GetOrder(token int) ([]order.HistoryOrder, error) {

	data, err := usecase.transactionData.SelectOrder(token)
	if err != nil {
		return nil, err
	}

	return data, nil
}
