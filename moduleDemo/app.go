package main

import (
	"fmt"
	"moduleDemo/utility"
)

func main() {
	num1, num2 := 10, 20.4
	fmt.Println(num1, num2)

	//constant
	fmt.Println("enter value")
	fmt.Scan(num1)
	fmt.Println(num1)

	fmt.Printf("values %s %v %v\n", "a", 1, 3.4)

	// var data = fmt.Sprint(num1)

	fmt.Println(string(num1))

	utility.Display()

	utility.Add5(&num1)

	fmt.Println(num1)
}
