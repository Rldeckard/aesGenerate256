package auth32

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func Random32bitString() string {
	b := make([]byte, 22)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func EncryptAES(key string, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher([]byte(key))
	fmt.Println("Checking cipher")
	if err != nil {
		log.Fatal("Failed to import decryption key.")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	// encrypt
	fmt.Println("Encrypting...")
	cipherText := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	// return hex string
	return hex.EncodeToString(cipherText)
}

func DecryptAES(key string, encodedText string) string {
	ciphertext, _ := hex.DecodeString(encodedText)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal("Failed to import decryption key.")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		log.Fatal("Failed to decrypt text. Check encryption key or regenerate credentials.")
	}
	return string(plaintext)
}
