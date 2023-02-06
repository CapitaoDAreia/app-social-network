package usersController

import (
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"net/http"
	"strings"
)

// Search for users in database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	usernameOrNickQuery := strings.ToLower(r.URL.Query().Get("user"))

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	userRepository := repository.NewUserRepository(DB)
	users, err := userRepository.SearchUsers(usernameOrNickQuery)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, users)
}
