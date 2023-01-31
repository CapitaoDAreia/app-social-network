package routes

import (
	"api-dvbk-socialNetwork/src/middlewares"
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
