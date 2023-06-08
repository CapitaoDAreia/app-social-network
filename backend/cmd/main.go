package main

import (
	"backend/internal/infraestructure/configuration"
	config "backend/internal/infraestructure/configuration"
	"backend/internal/infraestructure/database"
	"backend/internal/infraestructure/http/router/mux/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	configuration.GenerateSecretKey()
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

	returnR := routes.ConfigurateRoutes(r, DB)

	var PORT = fmt.Sprintf(":%v", config.APIPORT)

	fmt.Printf("Listening on PORT %v...\n", config.APIPORT)
	fmt.Println(MongoDB)
	log.Fatal(http.ListenAndServe(PORT, returnR))
}
