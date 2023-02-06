package loginController

import (
	"api-dvbk-socialNetwork/src/auth"
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/models"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"api-dvbk-socialNetwork/src/security"
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

	repository := repository.NewUserRepository(DB)
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
