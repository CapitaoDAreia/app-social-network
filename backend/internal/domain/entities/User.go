package entities

import (
	"backend/internal/infraestructure/http/auth"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents an user
type User struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Followers []string  `json:"followers,omitempty"`
	Following []string  `json:"following,omitempty"`
	Posts     []string  `json:"posts,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

type UserStageFlags struct {
	CanConsiderPasswordInValidateUser bool
	CanHashPassword                   bool
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

	if stage.CanConsiderPasswordInValidateUser && user.Password == "" {
		return errors.New("password is empty")
	}

	return nil
}

func (user *User) formatUserData(stage UserStageFlags) error {
	user.Username = strings.TrimSpace(user.Username)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if stage.CanHashPassword {
		hashedPassword, err := auth.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}

// Prepare user data to send for DB
func (user *User) PrepareUserData(stage UserStageFlags) error {
	if err := user.formatUserData(UserStageFlags{CanHashPassword: true}); err != nil {
		return err
	}

	if err := user.validateUserData(stage); err != nil {
		return err
	}

	return nil
}
