package main

import "fmt"

func main()  {
	fmt.Println("Hello arrays")
	var fruits[4] string;
	fruits[0] = "guava"
	fmt.Println(fruits)
	var anotherArray = [5]string{"a","b","c"}
	fmt.Println(anotherArray)
}