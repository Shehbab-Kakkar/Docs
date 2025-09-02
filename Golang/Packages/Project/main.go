package main

import (
	"fmt"

	"github.com/Shehbab-Kakkar/podcast/auth"
	"github.com/Shehbab-Kakkar/podcast/user"
)

func main() {
	auth.LoginWithCredentials("shehbab", "secret")
	session := auth.GetSession()
	fmt.Println("session", session)
	user := user.User{
		Email: "user@gmail.com",
		Name:  "John",
	}
	fmt.Println(user.Email, user.Name)
}
