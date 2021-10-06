package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	msg:= "Welcome to user inputs"
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any number to see the magic!")
	 
	// comma ok | error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}