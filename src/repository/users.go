package repository

import (
	"api-dvbk-socialNetwork/src/models"
	"database/sql"
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
	return 0, nil
}
