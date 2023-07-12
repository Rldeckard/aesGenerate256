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
		//needs a randomly generated 32 character string. Exactly 32 characters. The string is 22 characters, but it's encoded to 32. Confusing.
		key = auth32.Random32bitString()
	} else {
		key = prompt.Scan("Enter Key:")
	}
	// plaintext
	plainText := prompt.Scan("Enter string to encrypt: ")

	cipherText := auth32.EncryptAES(key, plainText)
	fmt.Println("Encrypted String: " + cipherText)
	fmt.Println("Decryption Key (do not lose): " + key)
	fmt.Println("\nTesting Secret.")
	providedKey := prompt.Scan("Enter Decryption Key (copy from above): ")
	fmt.Print("Is this your card? ")
	fmt.Println(auth32.DecryptAES(providedKey, cipherText))
	_ = prompt.Scan("Type stop or close window to end program after gathering information from above")

}
