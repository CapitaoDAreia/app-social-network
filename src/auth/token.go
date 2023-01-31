package auth

import (
	"api-dvbk-socialNetwork/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mockedKeySecret = []byte("SecretKey")

// Create an token that defines user permissions
func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["expt"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["userId"] = userID

	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return userToken.SignedString([]byte(config.SecretKey))
}
