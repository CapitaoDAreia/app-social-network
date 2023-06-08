package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringDatabaseKey = ""

	MongoConnectionString = ""

	DBPORT = ""

	MONGOPORT = ""

	APIPORT = ""

	err error

	SecretKey []byte

	User string

	Password string

	ServiceName string

	MongoServiceName string

	DBName string
)

// Iitialize ambient variables
func LoadAmbientConfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	ServiceName = os.Getenv("DB_SERVICE_NAME")
	APIPORT = os.Getenv("API_PORT")
	DBName = os.Getenv("DB_NAME_DATABASE")
	DBPORT = os.Getenv("DB_PORT")
	if err != nil {
		fmt.Printf("Error on .env PORT variable, assuming default PORT value: %v\n", DBPORT)
		DBPORT = "9000"
	}

	StringDatabaseKey = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		ServiceName,
		DBPORT,
		DBName,
	)

	// MONGODB
	MongoServiceName = os.Getenv("MONGO_SERVICE_NAME")
	MONGOPORT = os.Getenv("MONGO_PORT")
	if err != nil {
		fmt.Printf("Error on .env PORT variable, assuming default PORT value: %v\n", DBPORT)
		MONGOPORT = "27017"
	}

	MongoConnectionString = fmt.Sprintf(`mongodb://%s:%s@%s:%s/%s?authSource=admin`,
		User,
		Password,
		MongoServiceName,
		MONGOPORT,
		DBName,
	)

	SecretKey = []byte(os.Getenv("SecretKey"))

	fmt.Printf("SQLDatabaseKey: %v\n", StringDatabaseKey)
	fmt.Printf("MONGODatabaseKey: %v\n", MongoConnectionString)
}
