package mocks

import (
	"api-dvbk-socialNetwork/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (userRepository *UserRepositoryMock) CreateUser(user entities.User) (uint64, error) {
	args := userRepository.Called(user)

	return args.Get(0).(uint64), args.Error(1)
}
