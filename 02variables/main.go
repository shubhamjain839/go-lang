package main

import "fmt"

const token string = "supersecret" // By default public

func main()  {
	var username string = "Shubham"
	fmt.Println(username)
	fmt.Printf("Type: %T\n",username)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Type: %T\n",isTrue)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Type: %T\n",smallValue)

	var smallFloat float32 = 5435.3443434334343434343343
	fmt.Println(smallFloat)
	fmt.Printf("Type: %T\n",smallFloat)

	var justInt int 
	fmt.Println(justInt)
	fmt.Printf("Type: %T\n",justInt)

	var autoType = "jain"
	fmt.Println(autoType)
	fmt.Printf("Type: %T\n",autoType)

	noVarVariable:=1000000
	fmt.Println(noVarVariable)
	fmt.Printf("Type: %T\n",noVarVariable)

	fmt.Println(token)
	fmt.Printf("Type: %T\n",token)
}