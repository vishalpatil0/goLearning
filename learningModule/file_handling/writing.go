package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Writing to a file")
	var str string = "Hello, World!"
	os.WriteFile("balance.txt", []byte(str), 0644)
	fmt.Println("reading from a file")
	data, _ := os.ReadFile("balance.txt")
	str = string(data)
	fmt.Println("Reading from a file: ", str)
}
