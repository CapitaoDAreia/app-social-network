package routes

import (
	"backend/internal/infraestructure/http/middlewares"
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
func ConfigurateRoutes(r *mux.Router, db *sql.DB) *mux.Router {
	routes := []Route{}
	usersRoutes := ConfigUsersRoutes(db)
	postsRoutes := ConfigPostsRoutes(db)
	loginRoute := ConfigLoginRoutes(db)

	routes = append(routes, usersRoutes...)
	routes = append(routes, postsRoutes...)
	routes = append(routes, loginRoute)

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
