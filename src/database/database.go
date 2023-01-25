package database

import (
	config "api-dvbk-socialNetwork/src/Config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Opens connection with database and returns it.
func ConnectWithDatabase() (*sql.DB, error) {
	//Opens connection with database using key defined in config package.
	DB, err := sql.Open("mysql", config.StringDatabaseKey)
	if err != nil {
		return nil, err
	}

	//Close database connection if an error occur.
	if err = DB.Ping(); err != nil {
		DB.Close()
		return nil, err
	}

	return DB, nil
}
