package usersController

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	userSerialized, err := os.ReadFile("../../../../../test/resources/user.json")
	if err != nil {
		t.Fatal(err)
	}

	var user models.User
	if err := json.Unmarshal(userSerialized, &user); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                     string
		newUserBodyReq           *bytes.Buffer
		newUser                  models.User
		expectedStatusCode       int
		expectedCreateUserResult uint64
		expectedCreateUserError  error
	}{
		{
			name:                     "Success on createUser",
			newUserBodyReq:           bytes.NewBuffer(userSerialized),
			expectedStatusCode:       201,
			newUser:                  user,
			expectedCreateUserResult: 1,
			expectedCreateUserError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//httpconfig
			req := httptest.NewRequest(http.MethodPost, "/user", test.newUserBodyReq)
			res := httptest.NewRecorder()

			CreateUser(res, req)

			// asserts
			assert.Equal(t, res.Result().StatusCode, test.expectedStatusCode)

		})
	}
}
