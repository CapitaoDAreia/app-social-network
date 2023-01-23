package router

import (
	"github.com/gorilla/mux"
)

// Generates and return an mux router with routes setted
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
