package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func DeriveKey(secretKey string) string {
	hash := sha256.Sum256([]byte(secretKey))
	return hex.EncodeToString(hash[:])
}

func Encrypt(plaintext, hexSecretKey string) (string, error) {
	// 1. Take the hex version of hashed secret key and decode it into []byte
	key, err := hex.DecodeString(hexSecretKey)
	if err != nil {
		return "", err
	}
	// 2. Create a new block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// 3. Specify the mode of AES, in this case AES-GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// 4. Generate a cryptographically secure random nonce
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}
	// 5. Encrypt the plaintext
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	// 6. return hexadecimal string representation of ciphertext
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(ciphertextHex, hexSecretKey string) (string, error) {
	key, err := hex.DecodeString(hexSecretKey)
	if err != nil {
		return "", err
	}
	ciphertextBytes, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := ciphertextBytes[:gcm.NonceSize()]
	ciphertext := ciphertextBytes[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
