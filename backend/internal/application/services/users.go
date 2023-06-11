package services

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"fmt"
)

type UsersService interface {
	CreateUser(user entities.User) (string, error)
	SearchUsers(usernameOrNickQuery string) ([]entities.User, error)
	SearchUser(requestID string) (entities.User, error)
	UpdateUser(ID string, user entities.User) (uint64, error)
	DeleteUser(ID string) (uint64, error)
	SearchUserByEmail(email string) (entities.User, error)
	Follow(followedID, followerID string) error
	UnFollow(followedID, followerID string) error
	SearchFollowersOfAnUser(userID string) ([]string, error)
	SearchWhoAnUserFollow(userID string) ([]string, error)
	SearchUserPassword(userID string) (string, error)
	UpdateUserPassword(requestUserId string, hashedNewPasswordStringed string) (uint64, error)
}

type usersService struct {
	usersRepository repositories.UsersRepository
}

func NewUsersServices(userRepository repositories.UsersRepository) *usersService {
	return &usersService{
		userRepository,
	}
}

func (service *usersService) CreateUser(user entities.User) (string, error) {

	possibleExistentUser, _ := service.usersRepository.SearchUserByEmail(user.Email)

	if possibleExistentUser.Email == user.Email {
		return "", fmt.Errorf("This e-mail is already in use.")
	}

	if possibleExistentUser.Nick == user.Nick {
		return "", fmt.Errorf("This nick is already in use.")
	}

	if possibleExistentUser.Username == user.Username {
		return "", fmt.Errorf("This username is already in use.")
	}

	createdUserId, err := service.usersRepository.CreateUser(user)
	if err != nil {
		return "", err
	}
	return createdUserId, nil
}

func (service *usersService) SearchUsers(usernameOrNickQuery string) ([]entities.User, error) {

	allUsers, err := service.usersRepository.SearchUsers(usernameOrNickQuery)
	if err != nil {
		return []entities.User{}, err
	}

	return allUsers, nil
}

func (service *usersService) SearchUser(requestID string) (entities.User, error) {
	user, err := service.usersRepository.SearchUser(requestID)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (service *usersService) UpdateUser(ID string, user entities.User) (uint64, error) {

	// user, err := service.usersRepository.SearchUser(ID)
	// if err != nil {
	// 	return 0, fmt.Errorf("Error on SearchUser to UpdateUser: %s", err)
	// }

	modifiedCount, err := service.usersRepository.UpdateUser(ID, user)
	if err != nil {
		return modifiedCount, err
	}

	return modifiedCount, err
}

func (service *usersService) DeleteUser(ID string) (uint64, error) {
	deletedUserID, err := service.usersRepository.DeleteUser(ID)
	if err != nil {
		return 0, err
	}

	return deletedUserID, nil
}

func (service *usersService) SearchUserByEmail(email string) (entities.User, error) {
	user, err := service.usersRepository.SearchUserByEmail(email)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (service *usersService) Follow(followedID, followerID string) error {
	err := service.usersRepository.Follow(followedID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (service *usersService) UnFollow(followedID, followerID string) error {
	err := service.usersRepository.UnFollow(followedID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (service *usersService) SearchFollowersOfAnUser(userID string) ([]string, error) {
	followers, err := service.usersRepository.SearchFollowersOfAnUser(userID)
	if err != nil {
		return []string{}, err
	}

	return followers, nil
}

func (service *usersService) SearchWhoAnUserFollow(userID string) ([]string, error) {
	followedBy, err := service.usersRepository.SearchWhoAnUserFollow(userID)
	if err != nil {
		return []string{}, err
	}

	return followedBy, nil
}

func (service *usersService) SearchUserPassword(userID string) (string, error) {
	password, err := service.usersRepository.SearchUserPassword(userID)
	if err != nil {
		return "", err
	}

	return password, nil
}

func (service *usersService) UpdateUserPassword(requestUserId string, hashedNewPasswordStringed string) (uint64, error) {
	modifiedCount, err := service.usersRepository.UpdateUserPassword(requestUserId, hashedNewPasswordStringed)
	if err != nil {
		return 0, err
	}

	return modifiedCount, nil
}
