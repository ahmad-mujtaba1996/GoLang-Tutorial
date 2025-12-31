package main

//By putting go.mod file inside packages folder, inside import path we can remove 'packages/' part
import (
	"fmt"

	"github.com/ahmad-mujtaba1996/GoLang-Tutorial/auth"
	"github.com/ahmad-mujtaba1996/GoLang-Tutorial/user"

	// If you directly add third party package here without using 'go get' command, then you can still use it via 'go mod tidy' command which will add the package to go.mod and download it
	"github.com/fatih/color"
)

//packages/main.go:4:8: no required module provides package github.com/ahmad-mujtaba1996/GoLang-Tutorial/auth: go.mod file not found in current directory or any parent directory; see 'go help modules'
// To fix above error, we need to move into the packages folder and run 'go run main.go' command

func main() {

	// Getting error while calling the function from auth package. In order to fix it, we need to make the function
	auth.LoginWithCredentials("admin", "admin123")
	session := auth.GetSessionToken()
	fmt.Println("Session Token:", session)

	user := user.User{
		Name:  "Ahmad Mujtaba",
		Email: "ahmad@example.com"}

	color.Red(user.Name)
	color.Green(user.Email)

}
