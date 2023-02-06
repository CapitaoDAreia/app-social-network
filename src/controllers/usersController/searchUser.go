package usersController

import (
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Search for an specific user in database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	requestID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	userRepository := repository.NewUserRepository(DB)

	user, err := userRepository.SearchUser(requestID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, user)
}
