package service

import (
	"errors"
	"project/e-commerce/features/login"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.DataLoginInterface)
	t.Run("success Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(login.Core{Password: "$2a$10$uaFEnmPloNX9eoMbUNXb5eeq59BECM9hwi.pfSwB61rRfjijlM/N. "}, nil).Once()
		srv := New(repo)
		input := login.Core{Name: "Hery", Password: "asdf"}
		res, _, err := srv.Login(input)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Wrong username Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(login.Core{Password: "asgfasg"}, errors.New("username not found")).Once()
		srv := New(repo)
		input := login.Core{Name: "Hery", Password: "asdf"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.EqualError(t, err, "username not found")
		repo.AssertExpectations(t)
	})
}
