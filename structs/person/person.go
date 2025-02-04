package person

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	Person
}

func (user *Person) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	return &Admin{
		email: email,
		password: password,
		Person: Person{
			firstName: "Admin",
			lastName: "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}, nil
}

func New(firstName, lastName, birthDate string) (*Person, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &Person{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (user *Person) OutputUserDetails() {
	fmt.Println("First name: ",user.firstName)
	fmt.Println("Last name: ",user.lastName)
	fmt.Println("Birth date: ",user.birthDate)
	fmt.Println("Created at: ",user.createdAt.Format("02/01/2006"))
}