package postsController

import (
	"api-dvbk-socialNetwork/src/auth"
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// --
func DeletePost(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	postRequestID, err := strconv.ParseUint(params["postId"], 10, 64)
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

	currentPost, err := repository.SearchPost(postRequestID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if currentPost.AuthorID != tokenUserID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("Your are trying to delete something that is not yours!"))
		return
	}

	if err := repository.DeletePost(postRequestID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}
