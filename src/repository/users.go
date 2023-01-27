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
		return 0, err
	}
	defer statement.Close()

	execResult, err := statement.Exec(user.Username, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := execResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// Search for users by username or nick
func (u usersRepository) SearchUsers(usernameOrNickQuery string) ([]models.User, error) {
	usernameOrNickQuery = fmt.Sprintf("%%%s%%", usernameOrNickQuery) //%usernameOrNickQuery%

	rows, err := u.db.Query(
		"select id, username, nick, email, createdAt from users where username LIKE ? or nick LIKE ?",
		usernameOrNickQuery, usernameOrNickQuery,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u usersRepository) SearchUser(requestID uint64) (models.User, error) {
	rows, err := u.db.Query(
		"select id, username, nick, email, createdAt from users where id=?", requestID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
