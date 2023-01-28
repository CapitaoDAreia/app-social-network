package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

type UserStageFlags struct {
	ConsiderPassword bool
}

func (user *User) validateUserData(stage UserStageFlags) error {
	if user.Username == "" {
		return errors.New("username is empty")
	}

	if user.Nick == "" {
		return errors.New("nick is empty")
	}

	if user.Email == "" {
		return errors.New("email is empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email is invalid")
	}

	if stage.ConsiderPassword && user.Password == "" {
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
func (user *User) PrepareUserData(stage UserStageFlags) error {
	user.formatUserData()

	if err := user.validateUserData(stage); err != nil {
		return err
	}

	return nil
}
