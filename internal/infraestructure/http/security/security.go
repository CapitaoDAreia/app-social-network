package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	passwordSliceBytes := []byte(password)

	return bcrypt.GenerateFromPassword(passwordSliceBytes, bcrypt.DefaultCost)
}

func VerifyPassword(password, hashedPassword string) error {
	passwordSliceBytes := []byte(password)
	hashedPasswordSliceBytes := []byte(hashedPassword)

	return bcrypt.CompareHashAndPassword(hashedPasswordSliceBytes, passwordSliceBytes)
}
