package usersController

import (
	"api-dvbk-socialNetwork/internal/application/services/mocks"
	"api-dvbk-socialNetwork/internal/domain/entities"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	var user entities.User = entities.User{
		Nick:      "Admin",
		Username:  "Admin",
		Email:     "admin@email.com",
		Password:  "123456",
		CreatedAt: time.Now(),
	}

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(user)

	tests := []struct {
		name string
		user entities.User
	}{
		{
			name: "Success on CreateUser",
			user: user,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serviceMock := mocks.NewUsersServiceMock()
			serviceMock.On("CreateUser", user).Return(mock.AnythingOfType("uint64"), mock.AnythingOfType("error"))

			usersController := NewUsersController(serviceMock)

			rr := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/", nil)

			controller := http.HandlerFunc(usersController.CreateUser)
			controller.ServeHTTP(rr, req)
		})
	}
}
