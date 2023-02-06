package postsController

import (
	"api-dvbk-socialNetwork/src/auth"
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/models"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// --
func UpdatePost(w http.ResponseWriter, r *http.Request) {
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
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("Your are trying to update something that is not yours!"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	var updatedPost models.Post
	if err := json.Unmarshal(bodyRequest, &updatedPost); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if err := updatedPost.PreparePostData(); err != nil {
		responses.FormatResponseToCustomError(w, 500, errors.New("An error has occur when try to format Post data."))
		return
	}

	if err := repository.UpdatePost(postRequestID, updatedPost); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)

}
