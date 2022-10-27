package service

import (
	"errors"
	"project/e-commerce/features/user"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := new(mocks.DataInterface)
	input := user.Core{Name: "Hery", Email: "Hery@gmail.com", Password: "asdf", Phone: 123, Bio: "hallo aku senang", Saldo: 100000}

	t.Run("create success", func(t *testing.T) {

		repo.On("InsertData", mock.Anything).Return(1, nil).Once()

		useCase := New(repo)
		res, err := useCase.PostData(input)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})
	t.Run("error create data", func(t *testing.T) {

		repo.On("InsertData", mock.Anything).Return(-1, errors.New("there is some error")).Once()

		useCase := New(repo)
		res, err := useCase.PostData(input)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

	t.Run("empty data", func(t *testing.T) {

		input.Name = ""
		input.Password = ""
		input.Email = ""

		useCase := New(repo)
		res, err := useCase.PostData(input)
		assert.EqualError(t, err, "masukan nama ")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

func TestMyProfile(t *testing.T) {
	repo := new(mocks.DataInterface)
	data := user.Core{Name: "Hery", Email: "Hery@gmail.com", Phone: 123, Bio: "akun buat jajan ayang", Saldo: 90000000}

	t.Run("get data success", func(t *testing.T) {

		repo.On("SelectDataId", mock.Anything, mock.Anything).Return(data, nil).Once()

		useCase := New(repo)

		res, err := useCase.GetDataId(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, data, res)
		repo.AssertExpectations(t)
	})

	t.Run("error get data", func(t *testing.T) {

		repo.On("SelectDataId", mock.Anything, mock.Anything).Return(user.Core{}, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.GetDataId(1, 1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateProfile(t *testing.T) {
	repo := new(mocks.DataInterface)
	newData := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: 628999, Bio: "Akun foya foya", Saldo: 99000000000}

	t.Run("success update data", func(t *testing.T) {

		repo.On("UpdateData", mock.Anything).Return(-1, nil).Once()

		useCase := New(repo)

		res, err := useCase.PutDataId(newData)
		assert.NoError(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {

		newData.Password = "coco123"

		repo.On("UpdateData", mock.Anything).Return(-1, errors.New("there is some error")).Once()

		useCase := New(repo)

		res, _ := useCase.PutDataId(newData)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

}

func TestDeleteProfile(t *testing.T) {
	repo := new(mocks.DataInterface)

	t.Run("success delete data", func(t *testing.T) {

		repo.On("DeleteData", mock.Anything).Return(-1, nil).Once()

		useCase := New(repo)

		delete, err := useCase.DeleteAkun(1)
		assert.Nil(t, err)
		assert.Equal(t, -1, delete)
		repo.AssertExpectations(t)
	})

	t.Run("error delete data", func(t *testing.T) {

		repo.On("DeleteData", mock.Anything).Return(-1, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.DeleteAkun(1)
		assert.Error(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.DataInterface)
	data := user.Core{Name: "Hery", Email: "Hery@gmail.com", Phone: 123, Bio: "akun buat jajan ayang", Saldo: 90000000}

	t.Run("get data success", func(t *testing.T) {

		repo.On("MyProfile", mock.Anything).Return(data, nil).Once()

		useCase := New(repo)

		res, err := useCase.GetProfile(1)
		assert.NoError(t, err)
		assert.Equal(t, data, res)
		repo.AssertExpectations(t)
	})
	t.Run("error get data", func(t *testing.T) {

		repo.On("MyProfile", mock.Anything).Return(user.Core{}, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.GetProfile(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, res)
		repo.AssertExpectations(t)
	})
}
