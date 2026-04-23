//go:build ignore

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.Gen
	erateFromPassword([]byte("Admin@123456"), 10)
	fmt.Println(string(hash))
}
