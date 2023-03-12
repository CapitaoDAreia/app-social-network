package main

import (
	"api-dvbk-socialNetwork/internal/infraestructure/configuration"
	config "api-dvbk-socialNetwork/internal/infraestructure/configuration"
	"api-dvbk-socialNetwork/internal/infraestructure/database"
	router "api-dvbk-socialNetwork/internal/infraestructure/http/router/mux"
	"fmt"
	"log"
	"net/http"
)

func init() {
	configuration.GenerateSecretKey()
}

func main() {
	config.LoadAmbientConfig()
	fmt.Printf("PORT=%v\n", config.PORT)

	//Open connection with database
	DB, err := database.ConnectWithDatabase()
	if err != nil {
		panic(err)
	}

	//Generate routes to feed Server
	r := router.Generate(DB)

	//Generate PORT valur to feed Server
	var PORT = fmt.Sprintf(":%v", config.PORT)

	fmt.Printf("Listening on PORT %v...\n", config.PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
