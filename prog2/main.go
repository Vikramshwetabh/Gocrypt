package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Person struct matches the JSON structure
type Person struct {
	Name          string `json:"name"`
	Mobile        string `json:"mobile"`
	BloodGroup    string `json:"blood_group"`
	Email         string `json:"email"`
	AadhaarNumber string `json:"aadhaar_number"`
}

type Data struct {
	Person Person `json:"person"`
}

// Encrypt function
func encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) // Create a new AES cipher block
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data)) //created a slice for cipher text
	iv := ciphertext[:aes.BlockSize]                    // Generate a new IV
	_, err = io.ReadFull(rand.Reader, iv)               // Fill the IV with random data
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)           // Create a new CFB encrypter
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data) // Encrypt the data using XOR
	return ciphertext, nil
}

// Decrypt function
func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	data := ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}

func main() {
	// Read JSON file
	fileData, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Key for encryption (must be 16, 24, or 32 bytes)
	key := []byte("examplekey123456")

	// Encrypt the JSON data
	encrypted, err := encrypt(fileData, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Encrypted:", encrypted)

	// Decrypt the data
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	// Convert decrypted data back to struct
	var data Data
	err = json.Unmarshal(decrypted, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the result
	fmt.Println("Decrypted Data:", data.Person)
}
