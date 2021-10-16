package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else")

	var logincount int = 234
	var msg string
	if logincount > 233 {
		msg = "Login Failed"
	} else {
		msg = "Login Successful"
	}

	fmt.Println("Login Status: ", msg)

}
