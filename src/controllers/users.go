package controllers

import (
	"api-dvbk-socialNetwork/src/database"
	"api-dvbk-socialNetwork/src/models"
	"api-dvbk-socialNetwork/src/repository"
	"api-dvbk-socialNetwork/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	responses.FormatResponseToJSON(w, 200, user)
}

// Search for users in database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearcheUsers..."))
}

// Search for an specific user in database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearcheUsers..."))
}

// Update an user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser..."))
}

// Delete an user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser..."))
}
