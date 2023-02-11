package usersController

import (
	"api-dvbk-socialNetwork/internal/domain/entities"
	"api-dvbk-socialNetwork/internal/infraestructure/http/controllers/usersController/mocks"
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	json "encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {

	newUser := entities.User{
		Username: "username",
		Nick:     "username1",
		Email:    "email@email.com",
		Password: "123456",
	}
	newUserJSONConverted, _ := json.Marshal(newUser)

	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(newUserJSONConverted))

	tests := []struct {
		name                 string
		newUser              entities.User
		newUserJSONConverted []byte
		request              *http.Request
		requestError         error
		testResponseRecorder *httptest.ResponseRecorder
		createdUser          entities.CreatedUser

		createdUserReturn int
		createdUserError  error
	}{
		{
			name:                 "Success on create an user",
			newUser:              newUser,
			newUserJSONConverted: newUserJSONConverted,
			request:              request,
			requestError:         nil,
			testResponseRecorder: recorder,
			createdUser:          entities.CreatedUser{},
			createdUserReturn:    1,
			createdUserError:     nil,
		},
		{
			name:                 "...",
			newUser:              newUser,
			newUserJSONConverted: newUserJSONConverted,
			request:              request,
			requestError:         assert.AnError,
			testResponseRecorder: recorder,
			createdUser:          entities.CreatedUser{},
			createdUserReturn:    0,
			createdUserError:     assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responsesMock := mocks.NewResponsesMock()
			responsesMock.On("FormatResponseToCustomError", mock.AnythingOfType("http.ResponseWriter"), mock.AnythingOfType("int"), mock.AnythingOfType("error")).Return()
			responsesMock.On("FormatResponseToJSON", mock.AnythingOfType("http.ResponseWriter"), mock.AnythingOfType("int"), mock.AnythingOfType("error")).Return()

			DB := &sql.DB{}

			usersRepositoryMock := mocks.NewUserRepositoryMock(DB)
			usersRepositoryMock.On("CreateUser", mock.AnythingOfType("models.User")).Return(test.createdUserReturn, test.createdUserError)

			usersRepositoryMock.AssertNumberOfCalls(t, "CreateUser", 0)

			CreateUser(test.testResponseRecorder, test.request)
		})
	}
}
