package auth

import (
	config "backend/internal/infraestructure/configuration"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mockedKeySecret = []byte("SecretKey")

// Create an token that defines user permissions
func GenerateToken(userID string) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["expt"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["userId"] = userID

	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return userToken.SignedString([]byte(config.SecretKey))
}

// Verifies if token received in Request is valid
func ValidateToken(r *http.Request) error {

	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, tokenHasCorrespondentClaims := token.Claims.(jwt.MapClaims); tokenHasCorrespondentClaims && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}

func extractToken(r *http.Request) string {
	headerToken := r.Header.Get("Authorization")
	headerTokenSplited := strings.Split(headerToken, " ")

	if hasTwoValues := len(headerTokenSplited) == 2; !hasTwoValues {
		return " "
	}
	return headerTokenSplited[1]
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, tokenHasTrueValue := token.Method.(*jwt.SigningMethodHMAC); !tokenHasTrueValue {
		return nil, fmt.Errorf("Unexpected signature method: %v\n", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractUserID(r *http.Request) (string, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return "", err
	}

	if permissions, tokenHasCorrespondentClaims := token.Claims.(jwt.MapClaims); tokenHasCorrespondentClaims && token.Valid {
		userID := fmt.Sprintf("%s", permissions["userId"])
		if err != nil {
			return "", nil
		}

		return userID, nil
	}

	return "", errors.New("Invalid token")
}
