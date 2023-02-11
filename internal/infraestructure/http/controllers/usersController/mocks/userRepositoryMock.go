package mocks

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database/models"
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock(*sql.DB) *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (user *UserRepositoryMock) CreateUser(newUser models.User) (int, error) {
	args := user.Called(newUser)

	return args.Get(0).(int), args.Error(1)
}
