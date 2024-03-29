package models

import "time"

// User represents an user
type User struct {
	ID        string    `json:"_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Followers []string  `json:"followers,omitempty"`
	Following []string  `json:"following,omitempty"`
	Posts     []string  `json:"posts,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type UserStageFlags struct {
	CanConsiderPasswordInValidateUser bool
	CanHashPassword                   bool
}
