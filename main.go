package main

import "fmt"

//When
//1. When we need to update state

type User struct {
	name  string
	age   uint8
	email string
}

func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func Email(u *User) string {
	return u.email
}

func main() {
	user := User{
		email: "mark.teekens@outlook.com",
	}
	user.updateEmail("foobar@gmail.com")
	fmt.Println(user.Email())
}
