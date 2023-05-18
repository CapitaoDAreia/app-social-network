package configuration

import (
	"encoding/base64"
	"fmt"
	"os"
)

func GenerateSecretKey() {
	key := "my_secret_key"

	base64Result := base64.StdEncoding.EncodeToString([]byte(key))

	fmt.Println("SecretKey to userToken generated: " + base64Result)

	os.Setenv("SecretKey", base64Result)
}
