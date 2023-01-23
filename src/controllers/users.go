package controllers

import "net/http"

//Creates a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CriateUser..."))
}

//Search for users in database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearcheUsers..."))
}

// Search for an specific user in database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearcheUsers..."))
}

//Update an user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser..."))
}

//Delete an user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser..."))
}
