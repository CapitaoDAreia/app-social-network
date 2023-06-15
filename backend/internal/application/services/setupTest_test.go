package services

import (
	"backend/internal/domain/entities"
	"encoding/json"
	"os"
	"testing"
)

var User entities.User
var ExistentUser entities.User

func TestMain(m *testing.M) {
	userSerialized, _ := os.ReadFile("../../../test/resources/user.json")
	json.Unmarshal(userSerialized, &User)

	existentUserSerialized, _ := os.ReadFile("../../../test/resources/existent_user.json")
	json.Unmarshal(existentUserSerialized, &ExistentUser)

	os.Exit(m.Run())
}
