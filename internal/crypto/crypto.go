package crypto

import (
	"crypto/rand"
	"fmt"
)

func GenerateSecureToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
