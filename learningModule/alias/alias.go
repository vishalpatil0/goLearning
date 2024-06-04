package main

import "fmt"

type customString string

func (cstr customString) log() {
	fmt.Println(cstr)
}

func main() {
	var cstr customString = "Hello, World!"
	cstr.log()
}
