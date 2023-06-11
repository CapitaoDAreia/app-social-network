package services

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type UsersService interface {
	CreateUser(user entities.User) (string, error)
	SearchUsers(usernameOrNickQuery string) ([]entities.User, error)
	SearchUser(requestID uint64) (entities.User, error)
	UpdateUser(ID uint64, user entities.User) (uint64, error)
	DeleteUser(ID uint64) (uint64, error)
	SearchUserByEmail(email string) (entities.User, error)
	Follow(followedID, followerID uint64) error
	// UnFollow(followedID, followerID uint64) error
	// SearchFollowersOfAnUser(userID uint64) ([]entities.User, error)
	// SearchWhoAnUserFollow(userID uint64) ([]entities.User, error)
	// SearchUserPassword(userID uint64) (string, error)
	// UpdateUserPassword(requestUserId uint64, hashedNewPasswordStringed string) error
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
	createdUserId, err := service.usersRepository.CreateUser(user)
	if err != nil {
		return "0", err
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

func (service *usersService) SearchUser(requestID uint64) (entities.User, error) {
	user, err := service.usersRepository.SearchUser(requestID)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (service *usersService) UpdateUser(ID uint64, user entities.User) (uint64, error) {
	modifiedCount, err := service.usersRepository.UpdateUser(ID, user)
	if err != nil {
		return modifiedCount, err
	}

	return modifiedCount, err
}

func (service *usersService) DeleteUser(ID uint64) (uint64, error) {
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

func (service *usersService) Follow(followedID, followerID uint64) error {
	err := service.usersRepository.Follow(followedID, followerID)
	if err != nil {
		return err
	}

	return nil
}

// func (service *usersService) UnFollow(followedID, followerID uint64) error {
// 	err := service.usersRepository.UnFollow(followedID, followerID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (service *usersService) SearchFollowersOfAnUser(userID uint64) ([]entities.User, error) {
// 	followers, err := service.usersRepository.SearchFollowersOfAnUser(userID)
// 	if err != nil {
// 		return []entities.User{}, err
// 	}

// 	return followers, nil
// }

// func (service *usersService) SearchWhoAnUserFollow(userID uint64) ([]entities.User, error) {
// 	followedBy, err := service.usersRepository.SearchWhoAnUserFollow(userID)
// 	if err != nil {
// 		return []entities.User{}, err
// 	}

// 	return followedBy, nil
// }

// func (service *usersService) SearchUserPassword(userID uint64) (string, error) {
// 	password, err := service.usersRepository.SearchUserPassword(userID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return password, nil
// }

// func (service *usersService) UpdateUserPassword(requestUserId uint64, hashedNewPasswordStringed string) error {
// 	err := service.usersRepository.UpdateUserPassword(requestUserId, hashedNewPasswordStringed)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
