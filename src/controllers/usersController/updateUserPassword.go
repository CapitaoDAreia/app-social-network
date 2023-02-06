package usersController

import (
	"api-dvbk-socialNetwork/src/auth"
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/models"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"api-dvbk-socialNetwork/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	requestUserId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	tokenUserId, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if tokenUserId != requestUserId {
		responses.FormatResponseToCustomError(w, http.StatusUnauthorized, errors.New("Hmmm... Really?"))
		return
	}

	var password models.Password

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	if err := json.Unmarshal(bodyRequest, &password); err != nil {
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
	returnedPassword, err := repository.SearchUserPassword(requestUserId)

	if err := security.VerifyPassword(password.Current, returnedPassword); err != nil {
		responses.FormatResponseToCustomError(w, 500, errors.New("Current password not match!"))
		return
	}

	hashedNewPassword, err := security.Hash(password.New)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	hashedNewPasswordStringed := string(hashedNewPassword)

	if err := repository.UpdateUserPassword(requestUserId, hashedNewPasswordStringed); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}
