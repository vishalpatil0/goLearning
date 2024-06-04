package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.2))
	fmt.Println(add("a", "b"))
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

//any == interface{}
