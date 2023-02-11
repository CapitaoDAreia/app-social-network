package usersController

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	"api-dvbk-socialNetwork/internal/infraestructure/database/models"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	if err := user.PrepareUserData(models.UserStageFlags{CanConsiderPasswordInValidateUser: true}); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	userRepository := repository.NewUserRepository(DB)

	user.ID, err = userRepository.CreateUser(user)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 201, user)
}
