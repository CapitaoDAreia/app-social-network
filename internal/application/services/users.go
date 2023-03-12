package services

import (
	"api-dvbk-socialNetwork/internal/domain/entities"
	"api-dvbk-socialNetwork/internal/domain/repositories"
	"fmt"
)

type UsersService interface {
	CreateUser(user entities.User) error
}

type usersService struct {
	usersRepository repositories.UsersRepository
}

func NewUsersServices(userRepository repositories.UsersRepository) *usersService {
	return &usersService{
		userRepository,
	}
}

func (users *usersService) CreateUser(user entities.User) error {
	createdUserId, err := users.usersRepository.CreateUser(user)
	if err != nil {
		return err
	}
	fmt.Println(createdUserId)

	return nil
}
