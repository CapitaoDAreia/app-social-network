package auth

import "golang.org/x/crypto/bcrypt"

// Generate a hashed []byte from an string
func Hash(password string) ([]byte, error) {
	passwordSliceBytes := []byte(password)

	return bcrypt.GenerateFromPassword(passwordSliceBytes, bcrypt.DefaultCost)
}

// Receives an stringed password and an stringed hashedPassword and returns an error
func VerifyPassword(password, hashedPassword string) error {
	passwordSliceBytes := []byte(password)
	hashedPasswordSliceBytes := []byte(hashedPassword)

	return bcrypt.CompareHashAndPassword(hashedPasswordSliceBytes, passwordSliceBytes)
}
