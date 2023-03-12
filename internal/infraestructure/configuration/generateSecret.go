package configuration

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func GenerateSecretKey() {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	base64Result := base64.StdEncoding.EncodeToString([]byte(key))

	fmt.Println("SecretKey to userToken generated: " + base64Result)

	os.Setenv("SecretKey", base64Result)
}
