package usersController

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/responses"
	"net/http"
	"strings"
)

// Search for users in database
func GetUsers(w http.ResponseWriter, r *http.Request) {
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
