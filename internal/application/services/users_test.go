package services

import (
	"api-dvbk-socialNetwork/internal/domain/entities"
	"api-dvbk-socialNetwork/internal/infraestructure/database/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name                            string
		user                            entities.User
		expectedCreateUserReturn        uint64
		expectedCreateUserError         error
		expectedCreateUserNumberOfCalls int
	}{
		{
			name:                            "Success on CreateUser",
			user:                            User,
			expectedCreateUserReturn:        1,
			expectedCreateUserError:         nil,
			expectedCreateUserNumberOfCalls: 1,
		},
		{
			name:                            "Error on CreateUser",
			user:                            User,
			expectedCreateUserReturn:        0,
			expectedCreateUserError:         assert.AnError,
			expectedCreateUserNumberOfCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersRepositoryMock := mocks.NewUsersRepositoryMock()
			usersRepositoryMock.On("CreateUser", mock.AnythingOfType("entities.User")).Return(test.expectedCreateUserReturn, test.expectedCreateUserError)
			usersService := NewUsersServices(usersRepositoryMock)

			userID, err := usersService.CreateUser(test.user)

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
		requestID                       uint64
		expectedSearchUserReturn        entities.User
		expectedSearchUserError         error
		expectedSearchUserNumberOfCalls int
	}{
		{
			name:                            "Success on SearchUser",
			requestID:                       1,
			expectedSearchUserReturn:        entities.User{},
			expectedSearchUserError:         nil,
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
