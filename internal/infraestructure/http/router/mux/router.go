package router

import (
	"api-dvbk-socialNetwork/internal/infraestructure/http/router/mux/routes"
	"database/sql"

	"github.com/gorilla/mux"
)

// Generates and return an mux router with routes setted
func Generate(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	returnR := routes.Configurate(r, db)

	return returnR
}
