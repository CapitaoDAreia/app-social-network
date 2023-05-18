package services

import (
	"backend/internal/domain/entities"
	"encoding/json"
	"os"
	"testing"
)

var User entities.User

func TestMain(m *testing.M) {
	userSerialized, _ := os.ReadFile("../../../test/resources/user.json")
	json.Unmarshal(userSerialized, &User)

	os.Exit(m.Run())
}
