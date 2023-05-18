package usersController

import (
	"backend/internal/infraestructure/http/auth"
	"os"
	"testing"
)

var (
	ValidToken, DiffToken string
)

var Hashed string

func TestMain(m *testing.M) {
	ValidToken, _ = auth.GenerateToken(1)
	DiffToken, _ = auth.GenerateToken(2)
	hash, _ := auth.Hash("123456")

	Hashed = string(hash)

	os.Exit(m.Run())
}
