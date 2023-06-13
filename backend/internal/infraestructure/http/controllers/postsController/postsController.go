package postsController

import (
	"backend/internal/application/services"
	"backend/internal/domain/entities"
	"backend/internal/infraestructure/http/auth"
	"backend/internal/infraestructure/http/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type PostsController struct {
	postServices services.PostServices
}

func NewPostsController(postServices services.PostServices) *PostsController {
	return &PostsController{postServices}
}

// --
func (controller *PostsController) CreatePost(w http.ResponseWriter, r *http.Request) {
	userTokenId, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	var post entities.Post
	if err := json.Unmarshal(bodyRequest, &post); err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.AuthorID = userTokenId

	if err := post.PreparePostData(); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	post.ID, err = controller.postServices.CreatePost(post)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 201, post)

}

// --
func (controller *PostsController) DeletePost(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	postRequestID := params["postId"]

	currentPost, err := controller.postServices.SearchPost(postRequestID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if currentPost.AuthorID != tokenUserID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("Your are trying to delete something that is not yours!"))
		return
	}

	deletedCount, err := controller.postServices.DeletePost(postRequestID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusOK, deletedCount)
}

// --
func (controller *PostsController) GetPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID := parameters["postId"]

	post, err := controller.postServices.SearchPost(postID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, post)
}

// --
func (controller *PostsController) GetPosts(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnauthorized, err)
		return
	}

	posts, err := controller.postServices.SearchPosts(tokenUserID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, posts)
}

// --
func (controller *PostsController) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	requestUserId := parameters["userId"]

	userPosts, err := controller.postServices.SearchUserPosts(requestUserId)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, userPosts)
}

// --
func (controller *PostsController) LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID := parameters["postId"]

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnauthorized, err)
		return
	}

	if err := controller.postServices.LikePost(postID, tokenUserID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, nil)
}

// --
func (controller *PostsController) UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID := parameters["postId"]

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnauthorized, err)
		return
	}

	if err := controller.postServices.UnlikePost(postID, tokenUserID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, nil)
}

// --
func (controller *PostsController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	parameters := mux.Vars(r)

	postRequestID := parameters["postId"]

	currentPost, err := controller.postServices.SearchPost(postRequestID)
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

	var updatedPost entities.Post
	if err := json.Unmarshal(bodyRequest, &updatedPost); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if err := updatedPost.PreparePostData(); err != nil {
		responses.FormatResponseToCustomError(w, 500, errors.New("An error has occur when try to format Post data."))
		return
	}

	modifiedCount, err := controller.postServices.UpdatePost(postRequestID, updatedPost)

	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusOK, modifiedCount)

}
