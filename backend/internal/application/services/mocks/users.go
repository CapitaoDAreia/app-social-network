package mocks

import (
	"backend/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type UsersServiceMock struct {
	mock.Mock
}

func NewUsersServiceMock() *UsersServiceMock {
	return &UsersServiceMock{}
}

func (service *UsersServiceMock) CreateUser(user entities.User) (string, error) {
	args := service.Called(user)
	return args.Get(0).(string), args.Error(1)
}

func (service *UsersServiceMock) SearchUsers(usernameOrNickQuery string) ([]entities.User, error) {
	args := service.Called(usernameOrNickQuery)
	return args.Get(0).([]entities.User), args.Error(1)
}

func (service *UsersServiceMock) SearchUser(requestID string) (entities.User, error) {
	args := service.Called(requestID)
	return args.Get(0).(entities.User), args.Error(1)
}

func (service *UsersServiceMock) UpdateUser(ID string, user entities.User) (uint64, error) {
	args := service.Called(ID, user)
	return args.Get(0).(uint64), args.Error(1)
}

func (service *UsersServiceMock) DeleteUser(ID string) (uint64, error) {
	args := service.Called(ID)
	return args.Get(0).(uint64), args.Error(1)
}

func (service *UsersServiceMock) SearchUserByEmail(email string) (entities.User, error) {
	args := service.Called(email)
	return args.Get(0).(entities.User), args.Error(1)
}

func (service *UsersServiceMock) Follow(followedID, followerID string) error {
	args := service.Called(followedID, followerID)
	return args.Error(0)
}

func (service *UsersServiceMock) UnFollow(followedID, followerID string) error {
	args := service.Called(followedID, followerID)
	return args.Error(0)
}

func (service *UsersServiceMock) SearchFollowersOfAnUser(userID string) ([]string, error) {
	args := service.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func (service *UsersServiceMock) SearchWhoAnUserFollow(userID string) ([]string, error) {
	args := service.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func (service *UsersServiceMock) SearchUserPassword(userID string) (string, error) {
	args := service.Called(userID)
	return args.Get(0).(string), args.Error(1)
}

func (service *UsersServiceMock) UpdateUserPassword(requestUserId string, hashedNewPasswordStringed string) (uint64, error) {
	args := service.Called(requestUserId, hashedNewPasswordStringed)
	return args.Get(0).(uint64), args.Error(1)
}
