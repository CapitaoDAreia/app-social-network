package router

import (
	"api-dvbk-socialNetwork/src/router/routes"

	"github.com/gorilla/mux"
)

// Generates and return an mux router with routes setted
func Generate() *mux.Router {
	r := mux.NewRouter()
	returnR := routes.Configurate(r)

	return returnR
}
