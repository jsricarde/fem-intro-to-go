package main

import (
	"fmt"
)

//User type user
type User struct {
	Email string
}

func updateEmail(u *User) {
	u.Email = "jsricarde@gmail.com"
}

func main() {
	u := User{Email: "non@gmail.com"}
	fmt.Println("init", u)
	updateEmail(&u)
	fmt.Println("end", u)

}
