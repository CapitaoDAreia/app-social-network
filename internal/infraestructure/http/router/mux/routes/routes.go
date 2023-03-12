package routes

import (
	"api-dvbk-socialNetwork/internal/infraestructure/http/middlewares"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes is a struct that represents API routes
type Route struct {
	URI        string
	Method     string
	Controller func(http.ResponseWriter, *http.Request)
	NeedAuth   bool
}

// Config all routes in router
func Configurate(r *mux.Router, db *sql.DB) *mux.Router {
	routes := ConfigRoutes(db)
	routes = append(routes, LoginRoute)
	routes = append(routes, postRoutes...)

	// r.HandleFunc(route.URI, middlewares.Authenticate(route.Controller),).Methods(route.Method)

	for _, route := range routes {
		if route.NeedAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Controller),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Controller).Methods(route.Method)
		}
	}

	return r
}
