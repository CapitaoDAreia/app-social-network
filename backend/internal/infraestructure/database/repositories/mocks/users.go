package mocks

import (
	"backend/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

func NewUsersRepositoryMock() *UsersRepositoryMock {
	return &UsersRepositoryMock{}
}

func (usersRepository *UsersRepositoryMock) CreateUser(user entities.User) (string, error) {
	args := usersRepository.Called(user)

	return args.Get(0).(string), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchUsers(usernameOrNickQuery string) ([]entities.User, error) {
	args := usersRepository.Called(usernameOrNickQuery)

	return args.Get(0).([]entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchUser(requestID string) (entities.User, error) {
	args := usersRepository.Called(requestID)

	return args.Get(0).(entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) UpdateUser(ID string, user entities.User) (uint64, error) {
	args := usersRepository.Called(ID, user)

	return args.Get(0).(uint64), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) DeleteUser(ID string) (uint64, error) {
	args := usersRepository.Called(ID)

	return args.Get(0).(uint64), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchUserByEmail(email string) (entities.User, error) {
	args := usersRepository.Called(email)

	return args.Get(0).(entities.User), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) Follow(followedID, followerID string) error {
	args := usersRepository.Called(followedID, followerID)

	return args.Error(0)
}

func (usersRepository *UsersRepositoryMock) UnFollow(followedID, followerID string) error {
	args := usersRepository.Called(followedID, followerID)

	return args.Error(0)
}

func (usersRepository *UsersRepositoryMock) SearchFollowersOfAnUser(userID string) ([]string, error) {
	args := usersRepository.Called(userID)

	return args.Get(0).([]string), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchWhoAnUserFollow(userID string) ([]string, error) {
	args := usersRepository.Called(userID)

	return args.Get(0).([]string), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) SearchUserPassword(userID string) (string, error) {
	args := usersRepository.Called(userID)

	return args.Get(0).(string), args.Error(1)
}

func (usersRepository *UsersRepositoryMock) UpdateUserPassword(requestUserId string, hashedNewPasswordStringed string) (uint64, error) {
	args := usersRepository.Called(requestUserId, hashedNewPasswordStringed)

	return args.Get(0).(uint64), args.Error(1)
}
