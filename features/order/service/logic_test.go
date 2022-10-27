package service

import (
	"errors"
	"project/e-commerce/features/order"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := new(mocks.DataOrderInterface)
	address := order.AddressCore{Street: "jl. Margonda", City: "Depok", Province: "Jawa Barat", PostCode: 16519}
	payment := order.PaymentCore{Visa: "American Express", Name: "Hery", Number: 9912211, Cvv2: 9211, Month: 11, Year: 2027}

	t.Run("order success", func(t *testing.T) {

		repo.On("InsertData", mock.Anything, mock.Anything, mock.Anything).Return(1, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.PostData(1, address, payment)
		assert.Error(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

	t.Run("empty data", func(t *testing.T) {

		address.City = ""
		address.PostCode = 0
		address.Province = ""
		address.Street = ""
		address.TransactionID = 0

		payment.TransactionID = 0
		payment.Visa = ""
		payment.Name = ""
		payment.Number = 0
		payment.Cvv2 = 0
		payment.Month = 0
		payment.Year = 0

		useCase := New(repo)
		res, err := useCase.PostData(1, address, payment)
		assert.Error(t, err, "data tidak boleh kosong")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

func TestCancel(t *testing.T) {
	repo := new(mocks.DataOrderInterface)

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("CancelOrder", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteOrder(1, "waiting")
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("CancelOrder", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteOrder(1, "waiting")
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestConfirm(t *testing.T) {
	repo := new(mocks.DataOrderInterface)

	t.Run("Success Update data", func(t *testing.T) {
		repo.On("UpdateStatus", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutStatus(1, "waiting")
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("UpdateStatus", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.PutStatus(1, "waiting")
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetHistory(t *testing.T) {
	repo := new(mocks.DataOrderInterface)
	dataOrder := []order.HistoryOrder{{Images: "https://images.jpg", Name: "Sneakers", Price: 200000, Quantity: 2}}

	t.Run("Success Get order.", func(t *testing.T) {
		repo.On("SelectOrder", mock.Anything).Return(dataOrder, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetOrder(1)
		assert.NoError(t, err)
		assert.Equal(t, dataOrder, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get order.", func(t *testing.T) {
		repo.On("SelectOrder", mock.Anything).Return([]order.HistoryOrder{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetOrder(1)
		assert.Error(t, err)
		assert.Equal(t, []order.HistoryOrder([]order.HistoryOrder(nil)), result)
		repo.AssertExpectations(t)
	})
}
