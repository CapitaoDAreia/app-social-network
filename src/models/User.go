package models

import (
	"errors"
	"strings"
	"time"
)

// User represents an user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

func (user *User) validateUserData() error {
	if user.Username == "" {
		return errors.New("username is empty")
	}

	if user.Nick == "" {
		return errors.New("nick is empty")
	}

	if user.Email == "" {
		return errors.New("email is empty")
	}

	if user.Password == "" {
		return errors.New("password is empty")
	}

	return nil
}

func (user *User) formatUserData() {
	user.Username = strings.TrimSpace(user.Username)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}

// Prepare user data to send for DB
func (user *User) PrepareUserData() error {
	if err := user.validateUserData(); err != nil {
		return err
	}

	user.formatUserData()
	return nil
}
