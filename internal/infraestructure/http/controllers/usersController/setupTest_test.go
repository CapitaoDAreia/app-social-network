package usersController

import (
	"api-dvbk-socialNetwork/internal/infraestructure/http/auth"
	"api-dvbk-socialNetwork/internal/infraestructure/http/security"
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
	hash, _ := security.Hash("123456")

	Hashed = string(hash)

	os.Exit(m.Run())
}
