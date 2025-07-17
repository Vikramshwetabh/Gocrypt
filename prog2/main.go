package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"fmt"
	"io"
)

func main() {
	// Generate a strong, random key
	key := make([]byte, 32) // AES-256 key
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err.Error())
	}

	plaintext := []byte("This is the data for encryption.")

	//encryption

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil) // Seal encrypts and authenticates, appending the tag.
	fmt.Printf("Encrypted (hex): %x\n", ciphertext)
	fmt.Printf("Nonce (hex): %x\n", nonce)

	// --- Decryption ---

	// When decrypting, you need the same key and nonce.
	// You can extract the nonce from the beginning of the ciphertext if you appended it during encryption.
	// In this example, we kept the nonce separate for clarity.

	decryptedPlaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil) // Open decrypts and authenticates.
	if err != nil {
		panic(err.Error()) // Decryption failed, potentially due to incorrect key, nonce, or corrupted data.
	}

	fmt.Printf("Decrypted: %s\n", string(decryptedPlaintext))
}
