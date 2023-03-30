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
			repositoryMock := mocks.NewUsersRepositoryMock()
			repositoryMock.On("SearchUsers", test.nameOrNick).Return(test.expectedSearchUsersReturn, test.expectedSearchUsersError)
			services := NewUsersServices(repositoryMock)

			user, err := services.SearchUsers(test.nameOrNick)

			repositoryMock.AssertNumberOfCalls(t, "SearchUsers", test.expectedSearchUsersNumberOfCalls)
			assert.Equal(t, user, test.expectedSearchUsersReturn)
			assert.Equal(t, err, test.expectedSearchUsersError)
		})
	}
}
