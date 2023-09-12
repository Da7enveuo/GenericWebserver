package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

func generateSalt() string {
	// Generate a random slice of bytes
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err)
	}
	return string(salt)
}

func checkPassword(password string, user User) bool {
	// Hash the password with the salt
	hashedPassword := hashPassword(password, user.Salt)

	// Compare the hashed password with the stored hashed password
	return hashedPassword == user.HashedPassword
}

func hashPassword(password string, salt string) string {
	// Concatenate the password and salt
	saltedPassword := password + salt

	// Hash the salted password with the SHA256 algorithm
	hashedPassword := sha256.Sum256([]byte(saltedPassword))

	// Encode the hashed password as a base64 string
	encodedPassword := base64.StdEncoding.EncodeToString(hashedPassword[:])

	return encodedPassword
}

func createHashedPassword(password string) (string, string) {
	salt := generateSalt()
	hashedpw := hashPassword(password, salt)
	return hashedpw, salt
}
