package routes

import (
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
func Configurate(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, LoginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Controller).Methods(route.Method)
	}

	return r
}
