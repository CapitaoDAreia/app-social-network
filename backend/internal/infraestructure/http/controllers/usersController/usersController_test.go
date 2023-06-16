package usersController

import (
	"backend/internal/application/services/mocks"
	"backend/internal/domain/entities"
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
		expectedCreateUserResult string
		expectedStatusCode       int
		expectedErrorMessage     string
		responseIsAnError        bool
		expectedError            error
	}{
		{
			name:                     "Success on CreateUser",
			input:                    bytes.NewBuffer(userSerialized),
			expectedCreateUserResult: "1",
			expectedStatusCode:       http.StatusCreated,
			responseIsAnError:        false,
			expectedErrorMessage:     "",
			expectedError:            nil,
		},
		{
			name:                     "Error on CreateUser",
			input:                    bytes.NewBuffer(userSerialized),
			expectedCreateUserResult: "0",
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
					assert.Equal(t, fmt.Sprint(test.expectedCreateUserResult), strings.Trim(strings.TrimSpace(string(responseBody)), `"`))
				} else {
					assert.Equal(t, test.expectedErrorMessage, strings.TrimSpace(string(responseBody)))
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {

	tests := []struct {
		name                  string
		input                 string
		urlId                 string
		validToken            string
		userId                string
		expectedStatusCode    int
		expectedUpdatedReturn uint64
		expectedUpdatedError  error
	}{
		{
			name:                  "Success on UpdateUser",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            ValidToken,
			userId:                "1",
			expectedStatusCode:    204,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  nil,
		},
		{
			name:                  "Error on UpdateUser, unexistent url ID",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "",
			validToken:            ValidToken,
			userId:                "1",
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, ExtractUserID",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            ValidToken + "invalidate token",
			userId:                "1",
			expectedStatusCode:    401,
			expectedUpdatedReturn: 0,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, tokenId != requestId",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            DiffToken,
			userId:                "1",
			expectedStatusCode:    403,
			expectedUpdatedReturn: 0,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, empty bodyReq",
			input:                 "",
			urlId:                 "1",
			validToken:            ValidToken,
			userId:                "1",
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, broken bodyReq",
			input:                 `{"usernameupdated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            ValidToken,
			userId:                "1",
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on UpdateUser, incorrect field on bodyReq",
			input:                 `{"username":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            ValidToken,
			userId:                "1",
			expectedStatusCode:    500,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
		{
			name:                  "Error on call UpdateUser",
			input:                 `{"invalidField":"updated", "nick":"testupdated", "email":"user1@email.com"}`,
			urlId:                 "1",
			validToken:            ValidToken,
			userId:                "1",
			expectedStatusCode:    400,
			expectedUpdatedReturn: 1,
			expectedUpdatedError:  assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serviceMock := mocks.NewUsersServiceMock()
			serviceMock.On("UpdateUser", test.userId, mock.AnythingOfType("entities.User")).Return(test.expectedUpdatedReturn, test.expectedUpdatedError)
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
		input                    string
		expectedSearchUserReturn entities.User
		expectedSearchUserError  error
	}{
		{
			name:                     "Success on GetUser",
			requestID:                "1",
			expectedStatusCode:       200,
			input:                    "1",
			expectedSearchUserReturn: returnedUser,
			expectedSearchUserError:  nil,
		},
		{
			name:                     "Error on GetUser",
			requestID:                "1",
			expectedStatusCode:       500,
			input:                    "1",
			expectedSearchUserReturn: entities.User{},
			expectedSearchUserError:  assert.AnError,
		},
		{
			name:                     "Error on GetUser, empty requestId",
			requestID:                "",
			expectedStatusCode:       400,
			input:                    "1",
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
					ID:        "1",
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

func TestUpdateUserPassword(t *testing.T) {
	tests := []struct {
		name                     string
		validToken               string
		bodyReq                  string
		userId                   string
		searchUserPasswordReturn string
		searchUserPasswordError  error
		expectedUpdateUserError  error
		expectedUpdateUserResult uint64
		expectedStatusCode       int
	}{
		{
			name:                     "Success on UpdateUserPassword",
			validToken:               ValidToken,
			bodyReq:                  `{"current":"123456", "new":"789"}`,
			userId:                   "1",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  nil,
			expectedUpdateUserResult: 1,
			expectedStatusCode:       204,
		},
		{
			name:                     "Error on UpdateUserPassword",
			validToken:               ValidToken,
			bodyReq:                  `{"current":"123456", "new":"789"}`,
			userId:                   "1",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  assert.AnError,
			expectedUpdateUserResult: 0,
			expectedStatusCode:       500,
		},
		{
			name:                     "Error on UpdateUserPassword, empty data",
			validToken:               ValidToken,
			bodyReq:                  `{"current":"", "new":""}`,
			userId:                   "1",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  assert.AnError,
			expectedUpdateUserResult: 0,
			expectedStatusCode:       500,
		},
		{
			name:                     "Error on UpdateUserPassword, empty data",
			validToken:               ValidToken,
			bodyReq:                  ``,
			userId:                   "1",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  assert.AnError,
			expectedUpdateUserResult: 0,
			expectedStatusCode:       500,
		},
		{
			name:                     "Error on UpdateUserPassword, broken json",
			validToken:               ValidToken,
			bodyReq:                  `{"current"", "new":"}`,
			userId:                   "1",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  assert.AnError,
			expectedUpdateUserResult: 0,
			expectedStatusCode:       500,
		},
		{
			name:                     "Error on UpdateUserPassword, incorrect userId",
			validToken:               ValidToken,
			bodyReq:                  `{"current":"123456", "new":"789"}`,
			userId:                   "1000",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  assert.AnError,
			expectedUpdateUserResult: 0,
			expectedStatusCode:       401,
		},
		{
			name:                     "Error on UpdateUserPassword, invalid auth token",
			validToken:               ValidToken + "invalidate",
			bodyReq:                  `{"current":"123456", "new":"789"}`,
			userId:                   "1",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  assert.AnError,
			expectedUpdateUserResult: 0,
			expectedStatusCode:       500,
		},
		{
			name:                     "Error on UpdateUserPassword, empty userId",
			validToken:               ValidToken,
			bodyReq:                  `{"current":"123456", "new":"789"}`,
			userId:                   "",
			searchUserPasswordReturn: Hashed,
			searchUserPasswordError:  nil,
			expectedUpdateUserError:  nil,
			expectedUpdateUserResult: 0,
			expectedStatusCode:       400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			servicesMock := mocks.NewUsersServiceMock()
			servicesMock.On("SearchUserPassword", test.userId).Return(test.searchUserPasswordReturn, test.searchUserPasswordError)
			servicesMock.On("UpdateUserPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(test.expectedUpdateUserResult, test.expectedUpdateUserError)
			usersController := NewUsersController(servicesMock)
			req, _ := http.NewRequest("POST", "/", strings.NewReader(test.bodyReq))

			req.Header.Add("Authorization", "Bearer "+test.validToken)

			parameters := map[string]string{
				"userId": fmt.Sprintf("%s", test.userId),
			}
			req = mux.SetURLVars(req, parameters)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.UpdateUserPassword)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
		})
	}
}

func TestDeleteUser(t *testing.T) {

	tests := []struct {
		name                     string
		expectedStatusCode       int
		userId                   string
		validToken               string
		expectedDeleteUserError  error
		expectedDeleteUserReturn uint64
	}{
		{
			name:                     "Success on Delete",
			expectedStatusCode:       204,
			userId:                   "1",
			validToken:               ValidToken,
			expectedDeleteUserError:  nil,
			expectedDeleteUserReturn: 1,
		},
		{
			name:                     "Error on Delete",
			expectedStatusCode:       500,
			userId:                   "1",
			validToken:               ValidToken,
			expectedDeleteUserError:  assert.AnError,
			expectedDeleteUserReturn: 0,
		},
		{
			name:                     "Error on Delete, incorrect userId",
			expectedStatusCode:       401,
			userId:                   "122",
			validToken:               ValidToken,
			expectedDeleteUserError:  nil,
			expectedDeleteUserReturn: 0,
		},
		{
			name:                     "Error on Delete, invalid authToken",
			expectedStatusCode:       401,
			userId:                   "1",
			validToken:               ValidToken + "invalidate",
			expectedDeleteUserError:  nil,
			expectedDeleteUserReturn: 0,
		},
		{
			name:                     "Error on Delete, empty userId",
			expectedStatusCode:       400,
			userId:                   "",
			validToken:               ValidToken,
			expectedDeleteUserError:  nil,
			expectedDeleteUserReturn: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			servicesMock := mocks.NewUsersServiceMock()
			servicesMock.On("DeleteUser", mock.AnythingOfType("string")).Return(test.expectedDeleteUserReturn, test.expectedDeleteUserError)
			usersController := NewUsersController(servicesMock)

			req, _ := http.NewRequest("DELETE", "/users", nil)
			parameters := map[string]string{
				"userId": test.userId,
			}
			req = mux.SetURLVars(req, parameters)
			req.Header.Add("Authorization", "Bearer "+test.validToken)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.DeleteUser)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
		})
	}

}

func TestFollowUser(t *testing.T) {

	tests := []struct {
		name                string
		expectedStatusCode  int
		followedId          string
		validToken          string
		expectedFollowError error
	}{
		{
			name:                "Succcess on FollowUser",
			expectedStatusCode:  204,
			followedId:          "2",
			validToken:          ValidToken,
			expectedFollowError: nil,
		},
		{
			name:                "Error on FollowUser",
			expectedStatusCode:  500,
			followedId:          "2",
			validToken:          ValidToken,
			expectedFollowError: assert.AnError,
		},
		{
			name:                "Error on FollowUser, empty followedId",
			expectedStatusCode:  400,
			followedId:          "",
			validToken:          ValidToken,
			expectedFollowError: nil,
		},
		{
			name:                "Error on FollowUser, invalid authToken",
			expectedStatusCode:  401,
			followedId:          "2",
			validToken:          ValidToken + "invalidate",
			expectedFollowError: nil,
		},
		{
			name:                "Error on FollowUser, invalid followed quals follower",
			expectedStatusCode:  403,
			followedId:          "1",
			validToken:          ValidToken,
			expectedFollowError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			servicesMock := mocks.NewUsersServiceMock()
			servicesMock.On("Follow", mock.AnythingOfType("uint64"), mock.AnythingOfType("uint64")).Return(test.expectedFollowError)
			usersController := NewUsersController(servicesMock)

			req, _ := http.NewRequest("POST", "/users/1/follow", nil)
			req.Header.Add("Authorization", "Bearer "+test.validToken)
			parameters := map[string]string{
				"userId": test.followedId,
			}
			req = mux.SetURLVars(req, parameters)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.FollowUser)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)

		})
	}

}

func TestUnFollowUser(t *testing.T) {
	tests := []struct {
		name                   string
		expectedStatusCode     int
		followedId             string
		validToken             string
		expectedUnFollowResult error
	}{
		{
			name:                   "Succcess on UnFollowUser",
			expectedStatusCode:     204,
			followedId:             "2",
			validToken:             ValidToken,
			expectedUnFollowResult: nil,
		},
		{
			name:                   "Error on UnFollowUser",
			expectedStatusCode:     500,
			followedId:             "2",
			validToken:             ValidToken,
			expectedUnFollowResult: assert.AnError,
		},
		{
			name:                   "Error on UnFollowUser, invalid token",
			expectedStatusCode:     403,
			followedId:             "2",
			validToken:             ValidToken + "Invalidate",
			expectedUnFollowResult: nil,
		},
		{
			name:                   "Error on UnFollowUser, empty userId",
			expectedStatusCode:     400,
			followedId:             "",
			validToken:             ValidToken,
			expectedUnFollowResult: nil,
		},
		{
			name:                   "Error on UnFollowUser, wrong userId",
			expectedStatusCode:     403,
			followedId:             "1",
			validToken:             ValidToken,
			expectedUnFollowResult: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			servicesMock := mocks.NewUsersServiceMock()
			servicesMock.On("UnFollow", mock.AnythingOfType("uint64"), mock.AnythingOfType("uint64")).Return(test.expectedUnFollowResult)
			usersController := NewUsersController(servicesMock)

			req, _ := http.NewRequest("POST", "/users/{userId}/unfollow", nil)
			req.Header.Add("Authorization", "Bearer "+test.validToken)
			parameters := map[string]string{
				"userId": test.followedId,
			}
			req = mux.SetURLVars(req, parameters)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.UnFollowUser)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
		})
	}
}

func TestGetFollowersOfAnUser(t *testing.T) {
	tests := []struct {
		name                               string
		expectedStatusCode                 int
		userId                             string
		expectedGetFollowersOfAnUserError  error
		expectedGetFollowersOfAnUserResult []entities.User
	}{
		{
			name:                               "Success on GetFollowersOfAnUser",
			expectedStatusCode:                 200,
			userId:                             "1",
			expectedGetFollowersOfAnUserError:  nil,
			expectedGetFollowersOfAnUserResult: []entities.User{},
		},
		{
			name:                               "Error on GetFollowersOfAnUser",
			expectedStatusCode:                 500,
			userId:                             "1",
			expectedGetFollowersOfAnUserError:  assert.AnError,
			expectedGetFollowersOfAnUserResult: []entities.User{},
		},
		{
			name:                               "Error on GetFollowersOfAnUser, empty userId",
			expectedStatusCode:                 400,
			userId:                             "",
			expectedGetFollowersOfAnUserError:  nil,
			expectedGetFollowersOfAnUserResult: []entities.User{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			servicesMock := mocks.NewUsersServiceMock()
			servicesMock.On("SearchFollowersOfAnUser", mock.AnythingOfType("uint64")).Return(test.expectedGetFollowersOfAnUserResult, test.expectedGetFollowersOfAnUserError)
			usersController := NewUsersController(servicesMock)

			req, _ := http.NewRequest("GET", "/", nil)
			parameters := map[string]string{
				"userId": test.userId,
			}
			req = mux.SetURLVars(req, parameters)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.GetFollowersOfAnUser)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
		})
	}
}

func TestGetWhoAnUserFollow(t *testing.T) {
	tests := []struct {
		name                             string
		expectedStatusCode               int
		userId                           string
		expectedGetWhoAnUserFollowError  error
		expectedGetWhoAnUserFollowResult []entities.User
	}{
		{
			name:                             "Success on GetWhoAnUserFollow",
			expectedStatusCode:               200,
			userId:                           "1",
			expectedGetWhoAnUserFollowError:  nil,
			expectedGetWhoAnUserFollowResult: []entities.User{},
		},
		{
			name:                             "Error on GetWhoAnUserFollow",
			expectedStatusCode:               500,
			userId:                           "1",
			expectedGetWhoAnUserFollowError:  assert.AnError,
			expectedGetWhoAnUserFollowResult: []entities.User{},
		},
		{
			name:                             "Error on GetWhoAnUserFollow, empty userId",
			expectedStatusCode:               400,
			userId:                           "",
			expectedGetWhoAnUserFollowError:  nil,
			expectedGetWhoAnUserFollowResult: []entities.User{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			servicesMock := mocks.NewUsersServiceMock()
			servicesMock.On("SearchWhoAnUserFollow", mock.AnythingOfType("uint64")).Return(test.expectedGetWhoAnUserFollowResult, test.expectedGetWhoAnUserFollowError)
			usersController := NewUsersController(servicesMock)

			req, _ := http.NewRequest("GET", "/", nil)
			parameters := map[string]string{
				"userId": test.userId,
			}
			req = mux.SetURLVars(req, parameters)

			rr := httptest.NewRecorder()

			controller := http.HandlerFunc(usersController.GetWhoAnUserFollow)
			controller.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
		})
	}
}
