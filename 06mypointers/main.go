package main

import "fmt"

func main()  {
	fmt.Println("Wecome to pointers")

	myNumber := 1241
	var ptr *int = &myNumber

	fmt.Println("Value of the pointer is ", ptr)
}