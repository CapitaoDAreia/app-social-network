package usersController

import (
	"api-dvbk-socialNetwork/internal/application/services"
	"api-dvbk-socialNetwork/internal/domain/entities"
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	"api-dvbk-socialNetwork/internal/infraestructure/database/models"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/responses"
	"api-dvbk-socialNetwork/internal/infraestructure/http/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"api-dvbk-socialNetwork/internal/infraestructure/http/auth"
	"errors"
	"strconv"

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

// Creates an user in database
func (controller *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	//Catch bodyRequest
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Put bodyRequest into a user typed based on a model
	var user entities.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	if err := user.PrepareUserData(entities.UserStageFlags{CanConsiderPasswordInValidateUser: true}); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	//Open connection with database
	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	//Create a newUser repo feeding it with DB connection previously opened
	userRepository := repository.NewUserRepository(DB)

	//Use CreateUser, a method of usersRepository, to Create a newUser feedinf the method with the userReceived in bodyRequest.
	user.ID, err = userRepository.CreateUser(user)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 201, user)
}

// Update an user in database
func (controller *UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	requestID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
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
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	var user entities.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if err := user.PrepareUserData(entities.UserStageFlags{CanConsiderPasswordInValidateUser: false}); err != nil {
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

	err = repository.UpdateUser(requestID, user)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 204, nil)
}

// Search for an specific user in database
func (controller *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
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

// Search for users in database
func (controller *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {
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

func (controller *UsersController) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
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

// Delete an user in database
func (controller *UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	requestID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if tokenUserID != requestID {
		responses.FormatResponseToCustomError(w, http.StatusUnauthorized, errors.New("How dare, you?"))
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	repository := repository.NewUserRepository(DB)
	if err := repository.DeleteUser(requestID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 204, nil)
}

// Sets an user to follow another
func (controller *UsersController) FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	parameters := mux.Vars(r)
	followedID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if followedID == followerID {
		responses.FormatResponseToCustomError(w, 500, errors.New("Do you want to follow yourself? Pff! "))
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}
	defer DB.Close()

	repository := repository.NewUserRepository(DB)

	if err := repository.Follow(followedID, followerID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}

func (controller *UsersController) UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	parameters := mux.Vars(r)

	followedID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusBadRequest, err)
		return
	}

	if followedID == followerID {
		responses.FormatResponseToCustomError(w, http.StatusForbidden, errors.New("You are fated to follow yourself forever!"))
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	repository := repository.NewUserRepository(DB)

	if err := repository.UnFollow(followedID, followerID); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, http.StatusNoContent, nil)
}

func (controller *UsersController) GetFollowersOfAnUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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
	followers, err := repository.SearchFollowersOfnAnUser(userID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, followers)

}

func (controller *UsersController) GetWhoAnUserFollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	repository := repository.NewUserRepository(DB)
	followers, err := repository.SearchWhoAnUserFollow(userID)
	if err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	responses.FormatResponseToJSON(w, 200, followers)
}
