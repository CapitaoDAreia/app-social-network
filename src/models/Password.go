package models

type Password struct {
	Current uint64 `json:"current"`
	New     string `json:"New"`
}
