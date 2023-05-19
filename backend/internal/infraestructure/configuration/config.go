package configuration

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDatabaseKey = ""

	PORT = 0

	err error

	SecretKey []byte
)

// Iitialize ambient variables
func LoadAmbientConfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		fmt.Printf("Error on .env PORT variable, assuming default PORT value: %v\n", PORT)
		PORT = 9000
	}

	StringDatabaseKey = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SERVICE_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME_DATABASE"),
	)

	SecretKey = []byte(os.Getenv("SecretKey"))

	fmt.Printf("DatabaseKey: %v\n", StringDatabaseKey)
}
