package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error occured %w", err)
	}

	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password %w", err)
	}
	return nil
}


func main() {
	pass := "12345"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Printf("Logged in! with hash %s", hashedPass)
}