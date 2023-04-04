package main

import (
	"api-dvbk-socialNetwork/internal/infraestructure/configuration"
	config "api-dvbk-socialNetwork/internal/infraestructure/configuration"
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	"api-dvbk-socialNetwork/internal/infraestructure/http/router/mux/routes"
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
	fmt.Printf("PORT=%v\n", config.PORT)

	DB, err := database.ConnectWithDatabase()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	returnR := routes.ConfigurateRoutes(r, DB)

	var PORT = fmt.Sprintf(":%v", config.PORT)

	fmt.Printf("Listening on PORT %v...\n", config.PORT)
	log.Fatal(http.ListenAndServe(PORT, returnR))
}
