package usersController

import (
	"api-dvbk-socialNetwork/internal/application/services/mocks"
	"api-dvbk-socialNetwork/internal/domain/entities"
	"api-dvbk-socialNetwork/internal/infraestructure/http/auth"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	userSerialized, err := os.ReadFile("../../../../../test/resources/user.json")
	invalidUserSerialized, err := os.ReadFile("../../../../../test/resources/invalid_user.json")

	if err != nil {
		t.Errorf("json")
	}

	var user entities.User
	json.Unmarshal(userSerialized, &user)

	tests := []struct {
		name                     string
		input                    *bytes.Buffer
		expectedCreateUserResult uint64
		expectedStatusCode       int
		expectedErrorMessage     string
		responseIsAnError        bool
		expectedError            error
	}{
		{
			name:                     "Success on CreateUser",
			input:                    bytes.NewBuffer(userSerialized),
			expectedCreateUserResult: 1,
			expectedStatusCode:       http.StatusCreated,
			responseIsAnError:        false,
			expectedErrorMessage:     "",
			expectedError:            nil,
		},
		{
			name:                     "Error on CreateUser",
			input:                    bytes.NewBuffer(userSerialized),
			expectedCreateUserResult: 0,
			expectedStatusCode:       http.StatusBadRequest,
			responseIsAnError:        true,
			expectedErrorMessage:     "{\"error\":\"error ocurred\"}",
			expectedError:            errors.New("error ocurred"),
		},
		{
			name:                 "Error on CreateUser, empty input",
			input:                bytes.NewBuffer([]byte{}),
			expectedStatusCode:   http.StatusBadRequest,
			responseIsAnError:    true,
			expectedErrorMessage: "{\"error\":\"unexpected end of JSON input\"}",
			expectedError:        assert.AnError,
		},
		{
			name:                 "Error on CreateUser, invalid user data",
			input:                bytes.NewBuffer(invalidUserSerialized),
			expectedStatusCode:   http.StatusBadRequest,
			responseIsAnError:    true,
			expectedErrorMessage: "{\"error\":\"nick is empty\"}",
			expectedError:        assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serviceMock := mocks.NewUsersServiceMock()
			serviceMock.On("CreateUser", mock.AnythingOfType("entities.User")).Return(test.expectedCreateUserResult, test.expectedError)

			usersController := NewUsersController(serviceMock)

			req := httptest.NewRequest("POST", "/users", test.input)
			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.CreateUser)
			controller.ServeHTTP(rr, req)

			responseBody, _ := ioutil.ReadAll(rr.Result().Body)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Error on status code got %d; expected %d", rr.Result().StatusCode, test.expectedStatusCode)
			} else {
				if !test.responseIsAnError {
					assert.Equal(t, fmt.Sprint(test.expectedCreateUserResult), strings.TrimSpace(string(responseBody)))
				} else {
					assert.Equal(t, test.expectedErrorMessage, strings.TrimSpace(string(responseBody)))
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {

	validToken, err := auth.GenerateToken(1)
	if err != nil {
		t.Errorf("%s ", err)
	}
	diffToken, err := auth.GenerateToken(2)
	if err != nil {
		t.Errorf("%s ", err)
	}

	tests := []struct {
		name                  string
		input                 string
		urlId                 string
		validToken            string
		userId                uint64
		expectedStatusCode    int
		expectedUpdatedReturn uint64
		expectedUpdatedError  error
	}{
		{
			name:                  "Success on UpdateUser",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            validToken,
			userId:                1,
			expectedStatusCode:    204,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  nil,
		},
		{
			name:                  "Error on UpdateUser, unexistent url ID",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "",
			validToken:            validToken,
			userId:                1,
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, ExtractUserID",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            validToken + "invalidate token",
			userId:                1,
			expectedStatusCode:    401,
			expectedUpdatedReturn: 0,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, tokenId != requestId",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            diffToken,
			userId:                1,
			expectedStatusCode:    403,
			expectedUpdatedReturn: 0,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, empty bodyReq",
			input:                 "",
			urlId:                 "1",
			validToken:            validToken,
			userId:                1,
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, broken bodyReq",
			input:                 `{"usernameupdated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            validToken,
			userId:                1,
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, incorrect field on bodyReq",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            validToken,
			userId:                1,
			expectedStatusCode:    500,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on call UpdateUser",
			input:                 `{"invalidField":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            validToken,
			userId:                1,
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serviceMock := mocks.NewUsersServiceMock()
			serviceMock.On("UpdateUser", test.userId, mock.AnythingOfType("entities.User")).Return(test.expectedUpdatedError)
			usersController := NewUsersController(serviceMock)

			req, _ := http.NewRequest("PUT", "/", strings.NewReader(test.input))
			req.Header.Add("Authorization", "Bearer "+test.validToken)
			params := map[string]string{
				"userId": test.urlId,
			}
			req = mux.SetURLVars(req, params)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.UpdateUser)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)

		})
	}
}

func TestGetUser(t *testing.T) {

	var returnedUser entities.User
	userSerialized, _ := os.ReadFile("../../../../../test/resources/user.json")
	json.Unmarshal(userSerialized, &returnedUser)

	tests := []struct {
		name                     string
		requestID                string
		expectedStatusCode       int
		input                    uint64
		expectedSearchUserReturn entities.User
		expectedSearchUserError  error
	}{
		{
			name:                     "Success on GetUser",
			requestID:                "1",
			expectedStatusCode:       200,
			input:                    1,
			expectedSearchUserReturn: returnedUser,
			expectedSearchUserError:  nil,
		},
		{
			name:                     "Error on GetUser",
			requestID:                "1",
			expectedStatusCode:       500,
			input:                    1,
			expectedSearchUserReturn: entities.User{},
			expectedSearchUserError:  assert.AnError,
		},
		{
			name:                     "Error on GetUser, empty requestId",
			requestID:                "",
			expectedStatusCode:       400,
			input:                    1,
			expectedSearchUserReturn: entities.User{},
			expectedSearchUserError:  assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userServiceMock := mocks.NewUsersServiceMock()
			userServiceMock.On("SearchUser", test.input).Return(test.expectedSearchUserReturn, test.expectedSearchUserError)
			usersController := NewUsersController(userServiceMock)

			req, _ := http.NewRequest("GET", "/users/", nil)
			params := map[string]string{
				"userId": test.requestID,
			}
			req = mux.SetURLVars(req, params)
			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.GetUser)

			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
		})
	}
}

func TestGetUsers(t *testing.T) {

	tests := []struct {
		name                      string
		input                     string
		expectedSearchUsersReturn []entities.User
		expectedSearchUsersError  error
		expectedStatusCode        int
	}{
		{
			name:  "Success on GetUsers",
			input: "",
			expectedSearchUsersReturn: []entities.User{
				{
					ID:        1,
					Username:  "",
					Nick:      "",
					Email:     "",
					Password:  "",
					CreatedAt: time.Now(),
				},
			},
			expectedSearchUsersError: nil,
			expectedStatusCode:       200,
		},
		{
			name:                      "Error on GetUsers",
			input:                     "",
			expectedSearchUsersReturn: []entities.User{},
			expectedSearchUsersError:  assert.AnError,
			expectedStatusCode:        500,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userServiceMock := mocks.NewUsersServiceMock()
			userServiceMock.On("SearchUsers", test.input).Return(test.expectedSearchUsersReturn, test.expectedSearchUsersError)
			usersController := NewUsersController(userServiceMock)

			req, _ := http.NewRequest("GET", "/users", nil)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.GetUsers)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
		})
	}
}
