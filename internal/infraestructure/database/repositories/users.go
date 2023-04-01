package repository

import (
	"api-dvbk-socialNetwork/internal/domain/entities"
	"database/sql"
	"fmt"
)

type UsersRepository struct {
	db *sql.DB
}

// NewUserRepository Receives a database opened in controller and instances it in users struct.
func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

// CreateUser Creates a user on database.
// This is a method of users struct.
func (u UsersRepository) CreateUser(user entities.User) (uint64, error) {
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
func (u UsersRepository) SearchUsers(usernameOrNickQuery string) ([]entities.User, error) {
	usernameOrNickQuery = fmt.Sprintf("%%%s%%", usernameOrNickQuery) //%usernameOrNickQuery%

	rows, err := u.db.Query(
		"select id, username, nick, email, createdAt from users where username LIKE ? or nick LIKE ?",
		usernameOrNickQuery, usernameOrNickQuery,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

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

func (u UsersRepository) SearchUser(requestID uint64) (entities.User, error) {
	rows, err := u.db.Query(
		"select id, username, nick, email, createdAt from users where id=?", requestID,
	)
	if err != nil {
		return entities.User{}, err
	}
	defer rows.Close()

	var user entities.User
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return entities.User{}, err
		}
	}

	return user, nil
}

func (u UsersRepository) UpdateUser(ID uint64, user entities.User) error {
	statement, err := u.db.Prepare(
		"update users set username=?, nick=?, email=? where id=?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(
		user.Username,
		user.Nick,
		user.Email,
		ID,
	); err != nil {
		return err
	}

	return nil
}

func (u UsersRepository) DeleteUser(ID uint64) error {
	statement, err := u.db.Prepare("delete from users where id=?")
	if err != nil {
		return err
	}

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (u UsersRepository) SearchUserByEmail(email string) (entities.User, error) {
	row, err := u.db.Query("select id, password from users where email=?", email)
	if err != nil {
		return entities.User{}, err
	}
	defer row.Close()

	var user entities.User

	for row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return entities.User{}, err
		}
	}

	return user, nil
}

func (u UsersRepository) Follow(followedID, followerID uint64) error {
	statement, err := u.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(followedID, followerID); err != nil {
		return err
	}

	return nil
}

func (u UsersRepository) UnFollow(followedID, followerID uint64) error {
	statement, err := u.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}

	if _, err := statement.Exec(
		followedID,
		followerID,
	); err != nil {
		return err
	}

	return nil
}

func (u UsersRepository) SearchFollowersOfAnUser(userID uint64) ([]entities.User, error) {
	rows, err := u.db.Query(
		`select u.id, u.username, u.nick, u.email, u.createdAt 
		from users u inner join followers s 
		on u.id = s.follower_id where s.user_id = ?`, userID,
	)
	if err != nil {
		return []entities.User{}, err
	}
	defer rows.Close()

	var followers []entities.User

	for rows.Next() {
		var user entities.User

		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return []entities.User{}, err
		}

		followers = append(followers, user)
	}

	return followers, nil
}

func (u UsersRepository) SearchWhoAnUserFollow(userID uint64) ([]entities.User, error) {
	rows, err := u.db.Query(`
		select u.id, u.username, u.nick, u.email, u.createdAt
		from users u inner join followers s on u.id = s.user_id where s.follower_id = ?
	`, userID)
	if err != nil {
		return []entities.User{}, err
	}
	defer rows.Close()

	var followers []entities.User

	for rows.Next() {
		var user entities.User

		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return []entities.User{}, err
		}

		followers = append(followers, user)
	}

	return followers, nil
}

func (u UsersRepository) SearchUserPassword(userID uint64) (string, error) {
	rows, err := u.db.Query(`select password from users where id = ? `, userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var searchedUser entities.User

	for rows.Next() {
		if err := rows.Scan(
			&searchedUser.Password,
		); err != nil {
			return "", err
		}
	}
	return searchedUser.Password, err
}

func (u UsersRepository) UpdateUserPassword(requestUserId uint64, hashedNewPasswordStringed string) error {
	statement, err := u.db.Prepare(`update users set password = ? where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(hashedNewPasswordStringed, requestUserId); err != nil {
		return err
	}

	return nil
}
