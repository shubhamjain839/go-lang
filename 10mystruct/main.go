package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	shubham := User{"Shubham Jain", 24, true, "shubham.jain@block8.com"}
	fmt.Println("User : ", shubham)
}

type User struct {
	Name     string
	Age      int
	Verified bool
	Email    string
}
