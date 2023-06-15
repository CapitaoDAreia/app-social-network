package usersController

import (
	"backend/internal/application/services"
	"backend/internal/domain/entities"
	"backend/internal/infraestructure/database/models"
	"backend/internal/infraestructure/http/auth"
	"backend/internal/infraestructure/http/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type UsersController struct {
	userService services.UsersService
}

func NewUsersController(userService services.UsersService) *UsersController {
	return &UsersController{
		userService,
	}
}

// Calls services to create an user
func (controller *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user entities.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	if err := user.PrepareUserData(entities.UserStageFlags{CanConsiderPasswordInValidateUser: true}); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	createdUserID, err := controller.userService.CreateUser(user)
	if err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	responses.FormatResponseToJSON(w, 201, createdUserID)
}

// Update an user in database
func (controller *UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	requestID := parameters["userId"]

	if requestID == "" {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, fmt.Errorf("Give-me an user ID."))
		return
	}

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnauthorized, err)
		return
	}

	if requestID != tokenUserID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("Are you sure that is really you?"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	var user entities.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	if err := user.PrepareUserData(entities.UserStageFlags{CanConsiderPasswordInValidateUser: false}); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	_, err = controller.userService.UpdateUser(requestID, user)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 204, nil)
}

// Search for an specific user in database
func (controller *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	requestID := parameters["userId"]

	if requestID == "" {
		responses.FormatResponseToCustomError(w, 400, fmt.Errorf("Give-me an user ID."))
		return
	}

	user, err := controller.userService.SearchUser(requestID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, user)
}

// Search for users in database
func (controller *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	usernameOrNickQuery := strings.ToLower(r.URL.Query().Get("user"))

	users, err := controller.userService.SearchUsers(usernameOrNickQuery)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, users)
}

func (controller *UsersController) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	requestUserId := parameters["userId"]

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

	returnedPassword, err := controller.userService.SearchUserPassword(requestUserId)

	if err := auth.VerifyPassword(password.Current, returnedPassword); err != nil {
		responses.FormatResponseToCustomError(w, 500, errors.New("Current password not match!"))
		return
	}

	hashedNewPassword, err := auth.Hash(password.New)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	hashedNewPasswordStringed := string(hashedNewPassword)

	_, err = controller.userService.UpdateUserPassword(requestUserId, hashedNewPasswordStringed)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}

// Delete an user in database
func (controller *UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	requestID := parameters["userId"]

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 401, err)
		return
	}

	if tokenUserID != requestID {
		responses.FormatResponseToCustomError(w, http.StatusUnauthorized, errors.New("How dare, you?"))
		return
	}

	_, err = controller.userService.DeleteUser(requestID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 204, nil)
}

// Sets an user to follow another
func (controller *UsersController) FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 401, err)
		return
	}

	parameters := mux.Vars(r)

	followedID := parameters["userId"]

	if followedID == followerID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("Do you want to follow yourself? Pff! "))
		return
	}

	if err := controller.userService.Follow(followedID, followerID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}

func (controller *UsersController) UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 403, err)
		return
	}

	parameters := mux.Vars(r)

	followedID := parameters["userId"]

	if followedID == followerID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("Do you want to follow yourself? Pff! "))
		return
	}

	if followedID == followerID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("You are fated to follow yourself forever!"))
		return
	}

	if err := controller.userService.UnFollow(followedID, followerID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}

func (controller *UsersController) GetFollowersOfAnUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID := parameters["userId"]

	followers, err := controller.userService.SearchFollowersOfAnUser(userID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, followers)

}

func (controller *UsersController) GetWhoAnUserFollow(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID := parameters["userId"]

	followers, err := controller.userService.SearchWhoAnUserFollow(userID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, followers)
}
