package loginController

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	"api-dvbk-socialNetwork/internal/infraestructure/database/models"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/auth"
	"api-dvbk-socialNetwork/internal/infraestructure/http/responses"
	"api-dvbk-socialNetwork/internal/infraestructure/http/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	repository := repository.NewUsersRepository(DB)
	foundedUser, err := repository.SearchUserByEmail(user.Email)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if err := security.VerifyPassword(user.Password, foundedUser.Password); err != nil {
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
