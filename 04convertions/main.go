package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to converions")
	fmt.Println("Enter any number between 1 and 10")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for entering")

	newNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64) 
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newNum)
	}
}