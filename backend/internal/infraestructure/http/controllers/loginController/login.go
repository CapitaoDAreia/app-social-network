package loginController

import (
	"backend/internal/application/services"
	"backend/internal/domain/entities"
	"backend/internal/infraestructure/http/auth"
	"backend/internal/infraestructure/http/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoginController struct {
	usersServices services.UsersService
}

func NewLoginController(loginServices services.UsersService) *LoginController {
	return &LoginController{loginServices}
}

func (controller *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	var user entities.User

	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	foundedUser, err := controller.usersServices.SearchUserByEmail(user.Email)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	fmt.Println(foundedUser.ID)
	fmt.Println(foundedUser)

	if err := auth.VerifyPassword(user.Password, foundedUser.Password); err != nil {
		responses.FormatResponseToCustomError(w, 401, err)
		return
	}

	userToken, err := auth.GenerateToken(foundedUser.ID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, userToken)
}
