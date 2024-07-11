package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginPword := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword))
	if err != nil {
		fmt.Println("You can't Login")
		return
	}
	fmt.Println("You are logged in ")
}
