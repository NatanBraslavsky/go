package main

import (
	"fmt"

	"exemple.com/struct/user"
)



func main() {
	firstName := getUserData("Digite o primeiro nome: ")
	secoundName := getUserData("Digite seu sobrenome")
	birthDate := getUserData("Digite sua data de aniversario: ")

	var appUser *user.User

	appUser, err := user.NewUser(firstName, secoundName, birthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(text string) string {
	fmt.Println(text)
	var value string
	fmt.Scanln(&value)
	return value
} 