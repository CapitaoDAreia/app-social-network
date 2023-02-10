package services

import (
	"api-dvbk-socialNetwork/internal/domain/entities"
	"api-dvbk-socialNetwork/internal/domain/repositories"
	"fmt"
)

type UsersService interface {
	SaveUser(user entities.User) error
}

type usersService struct {
	usersRepository repositories.UsersRepository
}

func (users *usersService) SaveUser(user entities.User) error {
	createdUserId, err := users.usersRepository.CreateUser(user)
	if err != nil {
		return err
	}
	fmt.Println(createdUserId)

	return nil
}
