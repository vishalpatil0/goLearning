package main

import "fmt"

func main() {
	num := 1
	switch {
	case 1 != 1:
		fmt.Println("it is one")
	case 2 != 2:
		fmt.Println("it is 2")
	default:
		fmt.Println("default statement")
	}

	switch num {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("default")
	}
}
