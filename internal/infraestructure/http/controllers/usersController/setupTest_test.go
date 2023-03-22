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

var Hashed []byte

func TestMain(m *testing.M) {
	ValidToken, _ = auth.GenerateToken(1)
	DiffToken, _ = auth.GenerateToken(2)
	Hashed, _ = security.Hash("password")

	os.Exit(m.Run())
}
