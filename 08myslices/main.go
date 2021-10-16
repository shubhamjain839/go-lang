package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to slices")

	var slice = []string{}

	fmt.Printf("Type of slice is -> %T\n", slice)

	slice = append(slice, "Hi", "there", "how")

	fmt.Println(slice)

	newSlice := make([]int,4)

	newSlice[0] = 2123
	newSlice[1] = 3123
	newSlice[2] = 6123
	newSlice[3] = 5123

	fmt.Println(newSlice)

	newSlice = append(newSlice, 5456, 6754, 5645)

	fmt.Println(newSlice)

	sort.Ints(newSlice)

	fmt.Println("Sorted slice", newSlice)

	var courses = []string{"react" ,"solidity", "ethereum", "flutter", "go"}

	fmt.Println(courses)
	courses = append(courses[:2], courses[3:]... )
	fmt.Println("coures", courses)
}