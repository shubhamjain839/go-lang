package main

import "fmt"

func main()  {
	fmt.Println("Hello it's maps")

	myMap := make(map[string]string)

	myMap["M"] = "Matic"
	myMap["ETH"] = "Ethereum"
	myMap["BTC"] = "Bitcoin"
	myMap["ETC"] = "Ethereum Classic"

	fmt.Println("Map is ", myMap)

	for key, value := range myMap {
		fmt.Printf("For key %v value is : %v\n", key, value)
	}
} 