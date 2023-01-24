package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String that contains key values to establish connection with database.
	StringDatabaseKey = ""

	//Port where API will run.
	PORT = 0

	//Servs to handle errors while catch PORT value in .env
	err error
)

// Iitialize ambient variables
func LoadAmbientConfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	//Catch API port number from .env and convert to int type
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		fmt.Printf("Error on .env PORT variable, assuming default PORT value: %v\n", PORT)
		PORT = 9000
	}

	//Catch data from .env to compouse the key to connect with DB
	StringDatabaseKey = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&Loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME_DATABASE"),
	)

	fmt.Printf("DatabaseKey: %v\n", StringDatabaseKey)
}
