package user

import(
	"fmt"
	"errors"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u *User) OutputUser() {
	// ...
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearOutput() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "Admin",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}