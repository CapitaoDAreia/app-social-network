package main

import (
	"backend/internal/infraestructure/configuration"
	config "backend/internal/infraestructure/configuration"
	"backend/internal/infraestructure/database"
	"backend/internal/infraestructure/http/middlewares"
	routes_package "backend/internal/infraestructure/http/router/mux/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	configuration.GenerateSecretKey()
}

func configurateRoutes(r *mux.Router, db *sql.DB) *mux.Router {
	routes := []routes_package.Route{}

	usersRoutes := routes_package.ConfigUsersRoutes(db)
	postsRoutes := routes_package.ConfigPostsRoutes(db)
	loginRoute := routes_package.ConfigLoginRoutes(db)

	routes = append(routes, usersRoutes...)
	routes = append(routes, postsRoutes...)
	routes = append(routes, loginRoute)

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

func main() {
	config.LoadAmbientConfig()
	fmt.Printf("PORT=%v\n", config.APIPORT)

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		panic(err)
	}

	MongoDB, err := database.Connect()
	if err != nil {
		panic(fmt.Errorf("Error connecting on mongoDB: %s", err))
	}

	r := mux.NewRouter()

	returnR := configurateRoutes(r, DB)

	var PORT = fmt.Sprintf(":%v", config.APIPORT)

	fmt.Printf("Listening on PORT %v...\n", config.APIPORT)
	fmt.Println(MongoDB)
	log.Fatal(http.ListenAndServe(PORT, returnR))
}
