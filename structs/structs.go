package main

import (
	"fmt"

	"example.com/structs/person"
)

func main() {
	firstName := getUserData("Enter first name: ")
	lastName := getUserData("Enter last name: ")
	birthDate := getUserData("Enter birth date: ")

	user, err := person.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin, err := person.NewAdmin(
		"email@example.com",
		"teste@123",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	user.OutputUserDetails()
	user.ClearUserName()
	user.OutputUserDetails()
}

func getUserData(promptText string) string {
	var input string
	fmt.Print(promptText)
	fmt.Scanln(&input)
	return input
}