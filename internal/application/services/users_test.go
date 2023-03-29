package services

import (
	"api-dvbk-socialNetwork/internal/domain/entities"
	"api-dvbk-socialNetwork/internal/infraestructure/database/repositories/mocks"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {

	var user entities.User
	userSerialized, _ := os.ReadFile("../../../test/resources/user.json")
	json.Unmarshal(userSerialized, &user)

	tests := []struct {
		name                            string
		user                            entities.User
		expectedCreateUserReturn        uint64
		expectedCreateUserError         error
		expectedCreateUserNumberOfCalls int
	}{
		{
			name:                            "Success on CreateUser",
			user:                            user,
			expectedCreateUserReturn:        1,
			expectedCreateUserError:         nil,
			expectedCreateUserNumberOfCalls: 1,
		},
		{
			name:                            "Error on CreateUser",
			user:                            user,
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
