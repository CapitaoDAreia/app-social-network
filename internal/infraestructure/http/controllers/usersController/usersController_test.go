package usersController

import (
	"api-dvbk-socialNetwork/internal/application/services/mocks"
	"api-dvbk-socialNetwork/internal/domain/entities"
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
	user := entities.User{
		Username: "admin",
		Nick:     "admin123",
		Email:    "admin@admin.com",
	}
	userJson, _ := json.Marshal(user)

	tests := []struct {
		name string
	}{
		{
			name: "Success on Update User",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serviceMock := mocks.NewUsersServiceMock()
			serviceMock.On("CreateUser", mock.AnythingOfType("entities.User")).Return(1, nil)
			usersController := NewUsersController(serviceMock)

			req := httptest.NewRequest("PUT", "/users/1", bytes.NewBuffer(userJson))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer mockToken")
			parameters := req.URL.Query()
			fmt.Println(parameters)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.UpdateUser)
			controller.ServeHTTP(rr, req)

			if rr.Result().StatusCode != 204 {
				t.Errorf("Error status code; expected 204 got %d", rr.Result().StatusCode)
			}
		})
	}
}
