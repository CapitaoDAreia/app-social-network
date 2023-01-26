package repository

import (
	"api-dvbk-socialNetwork/src/models"
	"database/sql"
	"fmt"
)

type usersRepository struct {
	db *sql.DB
}

// NewUserRepository Receives a database opened in controller and instances it in users struct.
func NewUserRepository(db *sql.DB) *usersRepository {
	return &usersRepository{db}
}

// CreateUser Creates a user on database.
// This is a method of users struct.
func (u usersRepository) CreateUser(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into users (username, nick, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		fmt.Println("users.go 25")
		return 0, err
	}
	defer statement.Close()

	execResult, err := statement.Exec(user.Username, user.Nick, user.Email, user.Password)
	if err != nil {
		fmt.Println("users.go 32")

		return 0, err
	}

	lastInsertedID, err := execResult.LastInsertId()
	if err != nil {
		fmt.Println("users.go 29")

		return 0, err
	}

	return uint64(lastInsertedID), nil
}
