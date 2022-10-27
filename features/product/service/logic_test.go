package service

import (
	"errors"
	"project/e-commerce/features/product"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddProduct(t *testing.T) {
	repo := new(mocks.DataProductInterface)
	dataInput := product.Core{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Detail: "Size L, Original verified, Made in US", Stock: 12, UserID: 3}

	t.Run("Success Insert data.", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(product.Core{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Detail: "Size L, Original verified, Made in US", UserID: 3}, nil).Once()

		usecase := New(repo)
		result, err := usecase.Create(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Insert data.", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(product.Core{}, errors.New("error")).Once()

		usecase := New(repo)
		result, err := usecase.Create(dataInput)
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})

}

func TestGetAll(t *testing.T) {
	repo := new(mocks.DataProductInterface)
	dataProduct := []product.Core{{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Detail: "Size L, Original verified, Made in US", Stock: 12, UserID: 3}}

	t.Run("Success Get all data.", func(t *testing.T) {
		repo.On("GetAll", mock.Anything, mock.Anything).Return(dataProduct, nil).Once()

		usecase := New(repo)
		result, err := usecase.ShowAll()
		assert.NoError(t, err)
		assert.Equal(t, dataProduct, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get all data.", func(t *testing.T) {
		repo.On("GetAll", mock.Anything, mock.Anything).Return([]product.Core{}, errors.New("error")).Once()

		usecase := New(repo)
		result, err := usecase.ShowAll()
		assert.Error(t, err)
		assert.Equal(t, []product.Core([]product.Core(nil)), result)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.DataProductInterface)
	dataProduct := product.Core{ID: 1, Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Detail: "Size L, Original verified, Made in US", Stock: 12, UserID: 3}

	t.Run("Success Get data by Id.", func(t *testing.T) {
		repo.On("GetByID", mock.Anything).Return(dataProduct, nil).Once()

		usecase := New(repo)
		result, err := usecase.ShowByID(1)
		assert.NoError(t, err)
		assert.Equal(t, dataProduct, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data by Id.", func(t *testing.T) {
		repo.On("GetByID", mock.Anything).Return(product.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, _ := usecase.ShowByID(1)
		assert.Error(t, errors.New("Error"))
		assert.NotEqual(t, -1, result)
		repo.AssertExpectations(t)
	})
}
func TestUpdate(t *testing.T) {
	repo := new(mocks.DataProductInterface)
	newData := product.Core{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Detail: "Size L, Original verified, Made in US", Stock: 12, UserID: 3}

	t.Run("Success Update data", func(t *testing.T) {
		repo.On("Edit", mock.Anything, mock.Anything).Return(newData, nil).Once()

		usecase := New(repo)

		result, err := usecase.Update(newData, 1)
		assert.NoError(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("Edit", mock.Anything, mock.Anything).Return(product.Core{}, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.Update(newData, 1)
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
}

func TestMyProduct(t *testing.T) {
	repo := new(mocks.DataProductInterface)
	dataProduct := []product.Core{{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Detail: "Size L, Original verified, Made in US", UserID: 1, Stock: 3}}

	t.Run("Success Get my product.", func(t *testing.T) {
		repo.On("GetMy", mock.Anything).Return(dataProduct, nil).Once()

		usecase := New(repo)
		result, err := usecase.ShowMy(1)
		assert.NoError(t, err)
		assert.Equal(t, result[0].UserID, dataProduct[0].UserID)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get my product.", func(t *testing.T) {
		repo.On("GetMy", mock.Anything).Return([]product.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.ShowMy(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.DataProductInterface)

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("Remove", mock.Anything, mock.Anything).Return(product.Core{}, nil).Once()

		usecase := New(repo)

		result, err := usecase.Delete(8)
		assert.Nil(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("Remove", mock.Anything, mock.Anything).Return(product.Core{}, errors.New("Error")).Once()

		usecase := New(repo)

		result, _ := usecase.Delete(8)
		assert.Error(t, errors.New("Failed"))
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
}
