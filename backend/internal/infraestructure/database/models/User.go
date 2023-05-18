package models

import (
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

type UserStageFlags struct {
	CanConsiderPasswordInValidateUser bool
	CanHashPassword                   bool
}
