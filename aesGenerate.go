package main

import (
	"fmt"
	"strings"

	"github.com/Rldeckard/aesGenerate32/authGen"
	"github.com/Rldeckard/sshRunCMD/userPrompt"
)

func main() {
	var key string
	keyMoveInput := prompt.Scan("Would you like to provide a key?[y/n]: ")
	keyMove := strings.ToLower(keyMoveInput)
	if keyMove == "n" || keyMove == "no" {
		key = aes256.Random32ByteString()
	} else {
		key = prompt.Scan("Enter Key:")
	}
	// plaintext
	plainText := prompt.Scan("Enter string to encrypt: ")

	cipherText := aes256.Encrypt(key, plainText)
	fmt.Println("Encrypted String: " + cipherText)
	fmt.Println("Decryption Key (do not lose): " + key)
	fmt.Println("\nTesting Secret.")
	providedKey := prompt.Scan("Enter Decryption Key (copy from above): ")
	fmt.Print("Is this your card? ")
	fmt.Println(aes256.Decrypt(providedKey, cipherText))
	_ = prompt.Scan("Type stop or close window to end program after gathering information from above")

}
