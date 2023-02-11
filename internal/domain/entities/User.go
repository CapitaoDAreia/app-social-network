package entities

// User represents an user
type User struct {
	Username string `json:"username,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
