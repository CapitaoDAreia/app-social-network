package mocks

import (
	"api-dvbk-socialNetwork/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

func NewUsersRepositoryMock() *UsersRepositoryMock {
	return &UsersRepositoryMock{}
}

func (usersRepository *UsersRepositoryMock) CreateUser(user entities.User) (uint64, error) {
	args := usersRepository.Called(user)

	return args.Get(0).(uint64), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchUsers(usernameOrNickQuery string) ([]entities.User, error) {
	args := usersRepository.Called(usernameOrNickQuery)

	return args.Get(0).([]entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchUser(requestID uint64) (entities.User, error) {
	args := usersRepository.Called(requestID)

	return args.Get(0).(entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) UpdateUser(ID uint64, user entities.User) error {
	args := usersRepository.Called(ID, user)

	return args.Error(0)
}

func (usersRepository *UsersRepositoryMock) DeleteUser(ID uint64) error {
	args := usersRepository.Called(ID)

	return args.Error(0)
}

func (usersRepository *UsersRepositoryMock) SearchUserByEmail(email string) (entities.User, error) {
	args := usersRepository.Called(email)

	return args.Get(0).(entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) Follow(followedID, followerID uint64) error {
	args := usersRepository.Called(followedID, followerID)

	return args.Error(0)
}

func (usersRepository *UsersRepositoryMock) UnFollow(followedID, followerID uint64) error {
	args := usersRepository.Called(followedID, followerID)

	return args.Error(0)
}

func (usersRepository *UsersRepositoryMock) SearchFollowersOfAnUser(userID uint64) ([]entities.User, error) {
	args := usersRepository.Called(userID)

	return args.Get(0).([]entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchWhoAnUserFollow(userID uint64) ([]entities.User, error) {
	args := usersRepository.Called(userID)

	return args.Get(0).([]entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchUserPassword(userID uint64) (string, error) {
	args := usersRepository.Called(userID)

	return args.Get(0).(string), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) UpdateUserPassword(requestUserId uint64, hashedNewPasswordStringed string) error {
	args := usersRepository.Called(requestUserId, hashedNewPasswordStringed)

	return args.Error(0)
}
