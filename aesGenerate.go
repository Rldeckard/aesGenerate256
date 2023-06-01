package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var plainText string
	var providedKey string
	var keyMoveInput string
	var key string
	fmt.Print("Would you like to provide a key?[y/n]: ")
	fmt.Scan(&keyMoveInput)
	keyMove := strings.ToLower(keyMoveInput)
	if keyMove == "n" || keyMove == "no" {
		//needs a randomly generated 32 character string. Exactly 32 characters. The string is 22 characters, but it's encoded to 32. Confusing.
		key = generateRandomString(22)
	} else {
		fmt.Print("Enter Key: ")
		fmt.Scan(&key)
	}

	// plaintext
	fmt.Print("Enter string to encrypt: ")
	fmt.Scan(&plainText)

	cipherText := EncryptAES([]byte(key), plainText)
	fmt.Println("Encrypted String: " + cipherText)
	fmt.Println("Decryption Key (do not lose): " + key)
	fmt.Println("\nTesting Secret.")
	fmt.Print("Enter Decryption Key (copy from above): ")
	fmt.Scan(&providedKey)
	fmt.Print("Is this your card? ")
	fmt.Println(DecryptAES([]byte(providedKey), cipherText))
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	CheckError(err)
	return base64.StdEncoding.EncodeToString(b)
}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	fmt.Println("Checking cipher")
	CheckError(err)

	gcm, err := cipher.NewGCM(c)
	CheckError(err)

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	CheckError(err)

	// encrypt
	fmt.Println("Encrypting...")
	cipherText := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	// return hex string
	return hex.EncodeToString(cipherText)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	gcm, err := cipher.NewGCM(c)
	CheckError(err)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	return string(plaintext)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
