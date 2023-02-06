package postsController

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// --

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	requestUserId, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repository := repository.NewPostsRepository(DB)
	userPosts, err := repository.SearchUserPosts(requestUserId)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, userPosts)
}
