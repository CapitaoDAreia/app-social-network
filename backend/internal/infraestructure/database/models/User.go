package models

import "time"

// User represents an user
type User struct {
	ID        uint64    `json:"_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type UserStageFlags struct {
	CanConsiderPasswordInValidateUser bool
	CanHashPassword                   bool
}
