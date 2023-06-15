package services

import (
	"backend/internal/domain/entities"
	"backend/internal/infraestructure/database/repositories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name                                   string
		newUser                                entities.User
		existentUser                           entities.User
		expectedCreateUserReturn               string
		expectedCreateUserError                error
		expectedCreateUserNumberOfCalls        int
		expectedSearchUserByEmailReturn        entities.User
		expectedSearchUserByEmailError         error
		expectedSearchUserByEmailNumberOfCalls int
	}{
		{
			name:                                   "Success on CreateUser",
			newUser:                                User,
			existentUser:                           ExistentUser,
			expectedCreateUserReturn:               "1",
			expectedCreateUserError:                nil,
			expectedCreateUserNumberOfCalls:        1,
			expectedSearchUserByEmailReturn:        User,
			expectedSearchUserByEmailError:         nil,
			expectedSearchUserByEmailNumberOfCalls: 1,
		},
		{
			name:                                   "Error on CreateUser",
			newUser:                                User,
			existentUser:                           ExistentUser,
			expectedCreateUserReturn:               "",
			expectedCreateUserError:                assert.AnError,
			expectedCreateUserNumberOfCalls:        1,
			expectedSearchUserByEmailReturn:        User,
			expectedSearchUserByEmailError:         nil,
			expectedSearchUserByEmailNumberOfCalls: 1,
		},
		{
			name:                                   "Error on CreateUser, user email already exists",
			newUser:                                User,
			existentUser:                           User,
			expectedCreateUserReturn:               "",
			expectedCreateUserError:                errors.New("This e-mail is already in use."),
			expectedCreateUserNumberOfCalls:        0,
			expectedSearchUserByEmailReturn:        User,
			expectedSearchUserByEmailError:         nil,
			expectedSearchUserByEmailNumberOfCalls: 1,
		},
		{
			name:    "Error on CreateUser, user nick already exists",
			newUser: User,
			existentUser: entities.User{
				Nick: "Admin1",
			},
			expectedCreateUserReturn:               "",
			expectedCreateUserError:                errors.New("This nick is already in use."),
			expectedCreateUserNumberOfCalls:        0,
			expectedSearchUserByEmailReturn:        User,
			expectedSearchUserByEmailError:         nil,
			expectedSearchUserByEmailNumberOfCalls: 1,
		},
		{
			name:    "Error on CreateUser, username already exists",
			newUser: User,
			existentUser: entities.User{
				Username: "Admin",
			},
			expectedCreateUserReturn:               "",
			expectedCreateUserError:                errors.New("This username is already in use."),
			expectedCreateUserNumberOfCalls:        0,
			expectedSearchUserByEmailReturn:        User,
			expectedSearchUserByEmailError:         nil,
			expectedSearchUserByEmailNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("CreateUser", mock.AnythingOfType("entities.User")).Return(test.expectedCreateUserReturn, test.expectedCreateUserError)
			usersRepositoryMock.On("SearchUserByEmail", mock.AnythingOfType("string")).Return(test.existentUser, test.expectedSearchUserByEmailError)
			usersService := NewUsersServices(usersRepositoryMock)

			userID, err := usersService.CreateUser(test.newUser)

			usersRepositoryMock.AssertNumberOfCalls(t, "CreateUser", test.expectedCreateUserNumberOfCalls)
			assert.Equal(t, test.expectedCreateUserError, err)
			assert.Equal(t, test.expectedCreateUserReturn, userID)
		})
	}
}

func TestSearchUsers(t *testing.T) {
	tests := []struct {
		name                             string
		nameOrNick                       string
		expectedSearchUsersReturn        []entities.User
		expectedSearchUsersError         error
		expectedSearchUsersNumberOfCalls int
	}{
		{
			name:                             "Success on SearchUsers",
			nameOrNick:                       "Admin",
			expectedSearchUsersReturn:        []entities.User{User},
			expectedSearchUsersError:         nil,
			expectedSearchUsersNumberOfCalls: 1,
		},
		{
			name:                             "Error on SearchUsers",
			nameOrNick:                       "Admin",
			expectedSearchUsersReturn:        []entities.User{},
			expectedSearchUsersError:         assert.AnError,
			expectedSearchUsersNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("SearchUsers", test.nameOrNick).Return(test.expectedSearchUsersReturn, test.expectedSearchUsersError)
			services := NewUsersServices(usersRepositoryMock)

			users, err := services.SearchUsers(test.nameOrNick)

			usersRepositoryMock.AssertNumberOfCalls(t, "SearchUsers", test.expectedSearchUsersNumberOfCalls)
			assert.Equal(t, users, test.expectedSearchUsersReturn)
			assert.Equal(t, err, test.expectedSearchUsersError)
		})
	}
}

func TestSearchUser(t *testing.T) {
	tests := []struct {
		name                            string
		requestID                       string
		expectedSearchUserReturn        entities.User
		expectedSearchUserError         error
		expectedSearchUserNumberOfCalls int
	}{
		{
			name:                            "Success on SearchUser",
			requestID:                       "1",
			expectedSearchUserReturn:        User,
			expectedSearchUserError:         nil,
			expectedSearchUserNumberOfCalls: 1,
		},
		{
			name:                            "Error on SearchUser",
			requestID:                       "1",
			expectedSearchUserReturn:        entities.User{},
			expectedSearchUserError:         assert.AnError,
			expectedSearchUserNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("SearchUser", test.requestID).Return(test.expectedSearchUserReturn, test.expectedSearchUserError)
			services := NewUsersServices(usersRepositoryMock)

			user, err := services.SearchUser(test.requestID)

			usersRepositoryMock.AssertNumberOfCalls(t, "SearchUser", test.expectedSearchUserNumberOfCalls)
			assert.Equal(t, test.expectedSearchUserReturn, user)
			assert.Equal(t, test.expectedSearchUserError, err)

		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name                            string
		ID                              string
		user                            entities.User
		expectedUpdateUserReturn        error
		expectedUpdateUserNumberOfCalls int
	}{
		{
			name:                            "Success on UpdateUser",
			ID:                              "1",
			user:                            User,
			expectedUpdateUserReturn:        nil,
			expectedUpdateUserNumberOfCalls: 1,
		},
		{
			name:                            "Error on UpdateUser",
			ID:                              "1",
			user:                            entities.User{},
			expectedUpdateUserReturn:        assert.AnError,
			expectedUpdateUserNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("UpdateUser", test.ID, test.user).Return(test.expectedUpdateUserReturn)
			services := NewUsersServices(usersRepositoryMock)

			err, _ := services.UpdateUser(test.ID, test.user)

			usersRepositoryMock.AssertNumberOfCalls(t, "UpdateUser", test.expectedUpdateUserNumberOfCalls)
			assert.Equal(t, test.expectedUpdateUserReturn, err)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name                            string
		ID                              string
		expectedDeleteUserReturn        error
		expectedDeleteUserNumberOfCalls int
	}{
		{
			name:                            "Success on DeleteUser",
			ID:                              "1",
			expectedDeleteUserReturn:        nil,
			expectedDeleteUserNumberOfCalls: 1,
		},
		{
			name:                            "Error on DeleteUser",
			ID:                              "1",
			expectedDeleteUserReturn:        assert.AnError,
			expectedDeleteUserNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("DeleteUser", test.ID).Return(test.expectedDeleteUserReturn)
			services := NewUsersServices(usersRepositoryMock)

			err, _ := services.DeleteUser(test.ID)

			usersRepositoryMock.AssertNumberOfCalls(t, "DeleteUser", test.expectedDeleteUserNumberOfCalls)
			assert.Equal(t, test.expectedDeleteUserReturn, err)
		})
	}
}

func TestSearchUserByEmaill(t *testing.T) {
	tests := []struct {
		name                                    string
		userEmail                               string
		expectedSearchUserByEmaillReturn        entities.User
		expectedSearchUserByEmaillError         error
		expectedSearchUserByEmaillNumberOfCalls int
	}{
		{
			name:                                    "Success on SearchUserByEmaill",
			userEmail:                               "admin@admin.com",
			expectedSearchUserByEmaillReturn:        User,
			expectedSearchUserByEmaillError:         nil,
			expectedSearchUserByEmaillNumberOfCalls: 1,
		},
		{
			name:                                    "Error on SearchUserByEmaill",
			userEmail:                               "admin@admin.com",
			expectedSearchUserByEmaillReturn:        entities.User{},
			expectedSearchUserByEmaillError:         assert.AnError,
			expectedSearchUserByEmaillNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("SearchUserByEmail", test.userEmail).Return(test.expectedSearchUserByEmaillReturn, test.expectedSearchUserByEmaillError)

			services := NewUsersServices(usersRepositoryMock)

			user, err := services.SearchUserByEmail(test.userEmail)

			usersRepositoryMock.AssertNumberOfCalls(t, "SearchUserByEmail", test.expectedSearchUserByEmaillNumberOfCalls)
			assert.Equal(t, test.expectedSearchUserByEmaillReturn, user)
			assert.Equal(t, test.expectedSearchUserByEmaillError, err)
		})
	}
}

func TestFollow(t *testing.T) {
	tests := []struct {
		name                        string
		followedID                  string
		followerID                  string
		expectedFollowReturn        error
		expectedFollowNumberOfCalls int
	}{
		{
			name:                        "Success on Follow",
			followedID:                  "1",
			followerID:                  "2",
			expectedFollowReturn:        nil,
			expectedFollowNumberOfCalls: 1,
		},
		{
			name:                        "Error on Follow",
			followedID:                  "1",
			followerID:                  "2",
			expectedFollowReturn:        assert.AnError,
			expectedFollowNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("Follow", test.followedID, test.followerID).Return(test.expectedFollowReturn)

			services := NewUsersServices(usersRepositoryMock)

			err := services.Follow(test.followedID, test.followerID)

			usersRepositoryMock.AssertNumberOfCalls(t, "Follow", test.expectedFollowNumberOfCalls)
			assert.Equal(t, test.expectedFollowReturn, err)
		})
	}
}

func TestUnFollow(t *testing.T) {
	tests := []struct {
		name                          string
		followedID                    string
		followerID                    string
		expectedUnFollowReturn        error
		expectedUnFollowNumberOfCalls int
	}{
		{
			name:                          "Success on UnFollow",
			followedID:                    "1",
			followerID:                    "2",
			expectedUnFollowReturn:        nil,
			expectedUnFollowNumberOfCalls: 1,
		},
		{
			name:                          "Error on UnFollow",
			followedID:                    "1",
			followerID:                    "2",
			expectedUnFollowReturn:        assert.AnError,
			expectedUnFollowNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("UnFollow", test.followedID, test.followerID).Return(test.expectedUnFollowReturn)

			services := NewUsersServices(usersRepositoryMock)

			err := services.UnFollow(test.followedID, test.followerID)

			usersRepositoryMock.AssertNumberOfCalls(t, "UnFollow", test.expectedUnFollowNumberOfCalls)
			assert.Equal(t, test.expectedUnFollowReturn, err)
		})
	}
}

func TestSearchFollowersOfAnUser(t *testing.T) {
	tests := []struct {
		name                                         string
		userId                                       string
		expectedSearchFollowersOfAnUserError         error
		expectedSearchFollowersOfAnUserReturn        []entities.User
		expectedSearchFollowersOfAnUserNumberOfCalls int
	}{
		{
			name:                                  "Success on SearchFollowersOfAnUser",
			userId:                                "1",
			expectedSearchFollowersOfAnUserError:  nil,
			expectedSearchFollowersOfAnUserReturn: []entities.User{User},
			expectedSearchFollowersOfAnUserNumberOfCalls: 1,
		},
		{
			name:                                  "Error on SearchFollowersOfAnUser",
			userId:                                "1",
			expectedSearchFollowersOfAnUserError:  assert.AnError,
			expectedSearchFollowersOfAnUserReturn: []entities.User{},
			expectedSearchFollowersOfAnUserNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("SearchFollowersOfAnUser", test.userId).Return(test.expectedSearchFollowersOfAnUserReturn, test.expectedSearchFollowersOfAnUserError)

			services := NewUsersServices(usersRepositoryMock)

			users, err := services.SearchFollowersOfAnUser(test.userId)

			usersRepositoryMock.AssertNumberOfCalls(t, "SearchFollowersOfAnUser", test.expectedSearchFollowersOfAnUserNumberOfCalls)
			assert.Equal(t, test.expectedSearchFollowersOfAnUserReturn, users)
			assert.Equal(t, test.expectedSearchFollowersOfAnUserError, err)
		})
	}
}

func TestSearchWhoAnUserFollow(t *testing.T) {
	tests := []struct {
		name                                       string
		userId                                     string
		expectedSearchWhoAnUserFollowError         error
		expectedSearchWhoAnUserFollowReturn        []entities.User
		expectedSearchWhoAnUserFollowNumberOfCalls int
	}{
		{
			name:                                "Success on SearchWhoAnUserFollow",
			userId:                              "1",
			expectedSearchWhoAnUserFollowError:  nil,
			expectedSearchWhoAnUserFollowReturn: []entities.User{User},
			expectedSearchWhoAnUserFollowNumberOfCalls: 1,
		},
		{
			name:                                "Error on SearchWhoAnUserFollow",
			userId:                              "1",
			expectedSearchWhoAnUserFollowError:  assert.AnError,
			expectedSearchWhoAnUserFollowReturn: []entities.User{},
			expectedSearchWhoAnUserFollowNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("SearchWhoAnUserFollow", test.userId).Return(test.expectedSearchWhoAnUserFollowReturn, test.expectedSearchWhoAnUserFollowError)

			services := NewUsersServices(usersRepositoryMock)

			users, err := services.SearchWhoAnUserFollow(test.userId)

			usersRepositoryMock.AssertNumberOfCalls(t, "SearchWhoAnUserFollow", test.expectedSearchWhoAnUserFollowNumberOfCalls)
			assert.Equal(t, test.expectedSearchWhoAnUserFollowReturn, users)
			assert.Equal(t, test.expectedSearchWhoAnUserFollowError, err)
		})
	}
}

func TestSearchUserPassword(t *testing.T) {
	tests := []struct {
		name                                    string
		userId                                  string
		expectedSearchUserPasswordError         error
		expectedSearchUserPasswordReturn        string
		expectedSearchUserPasswordNumberOfCalls int
	}{
		{
			name:                                    "Success on SearchUserPassword",
			userId:                                  "1",
			expectedSearchUserPasswordError:         nil,
			expectedSearchUserPasswordReturn:        "",
			expectedSearchUserPasswordNumberOfCalls: 1,
		},
		{
			name:                                    "Error on SearchUserPassword",
			userId:                                  "1",
			expectedSearchUserPasswordError:         assert.AnError,
			expectedSearchUserPasswordReturn:        "",
			expectedSearchUserPasswordNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("SearchUserPassword", test.userId).Return(test.expectedSearchUserPasswordReturn, test.expectedSearchUserPasswordError)

			services := NewUsersServices(usersRepositoryMock)

			users, err := services.SearchUserPassword(test.userId)

			usersRepositoryMock.AssertNumberOfCalls(t, "SearchUserPassword", test.expectedSearchUserPasswordNumberOfCalls)
			assert.Equal(t, test.expectedSearchUserPasswordReturn, users)
			assert.Equal(t, test.expectedSearchUserPasswordError, err)
		})
	}
}

func TestUpdateUserPassword(t *testing.T) {
	tests := []struct {
		name                                    string
		requestUserId                           string
		hashedNewPasswordStringed               string
		expectedUpdateUserPasswordReturn        error
		expectedUpdateUserPasswordNumberOfCalls int
	}{
		{
			name:                                    "Success on UpdateUserPassword",
			requestUserId:                           "1",
			hashedNewPasswordStringed:               "",
			expectedUpdateUserPasswordReturn:        nil,
			expectedUpdateUserPasswordNumberOfCalls: 1,
		},
		{
			name:                                    "Error on UpdateUserPassword",
			requestUserId:                           "1",
			hashedNewPasswordStringed:               "",
			expectedUpdateUserPasswordReturn:        assert.AnError,
			expectedUpdateUserPasswordNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("UpdateUserPassword", test.requestUserId, test.hashedNewPasswordStringed).Return(test.expectedUpdateUserPasswordReturn)

			services := NewUsersServices(usersRepositoryMock)

			err, _ := services.UpdateUserPassword(test.requestUserId, test.hashedNewPasswordStringed)

			usersRepositoryMock.AssertNumberOfCalls(t, "UpdateUserPassword", test.expectedUpdateUserPasswordNumberOfCalls)
			assert.Equal(t, test.expectedUpdateUserPasswordReturn, err)
		})
	}
}
