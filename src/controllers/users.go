package controllers

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
	"strings"

	"github.com/gorilla/mux"
)

// Creates a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Catch bodyRequest
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.FormatResponseToCustomError(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Put bodyRequest into a user typed based on a model
	var user models.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.FormatResponseToCustomError(w, 400, err)
		return
	}

	if err := user.PrepareUserData(models.UserStageFlags{CanConsiderPasswordInValidateUser: true}); err != nil {
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

// Update an user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	var user models.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.FormatResponseToCustomError(w, 500, err)
		return
	}

	if err := user.PrepareUserData(models.UserStageFlags{CanConsiderPasswordInValidateUser: false}); err != nil {
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

// Delete an user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

func FollowUser(w http.ResponseWriter, r *http.Request) {
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

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
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
