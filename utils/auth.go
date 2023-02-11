package utils

import (
	"fmt"
	"strings"
)

func VerifyPassword(hashedPassword string, password string) bool {
	// err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// fmt.Println(err)
	fmt.Println(hashedPassword)
	return strings.Compare(hashedPassword, password) == 0
}

func HashPassword(password string) (string, error) {
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// return string(bytes), err
	return password, nil
}
