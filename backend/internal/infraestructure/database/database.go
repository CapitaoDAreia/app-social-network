package database

import (
	"backend/internal/infraestructure/configuration"
	config "backend/internal/infraestructure/configuration"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectWithDatabase() (*sql.DB, error) {
	DB, err := sql.Open("mysql", config.StringDatabaseKey)
	if err != nil {
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		DB.Close()
		return nil, err
	}

	return DB, nil
}

func Connect() (*mongo.Database, error) {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(configuration.MongoConnectionString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to mongoDB: %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("Error communicating to mongoDB: %s", err)
	}

	return client.Database(configuration.DBName), nil
}
