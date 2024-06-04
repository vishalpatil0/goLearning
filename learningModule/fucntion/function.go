package main

import (
	"errors"
	"fmt"
)

func main() {
	var name string = "data"
	value := demo(name)
	fmt.Println(value)

	division, remainder, err := intDivision(10, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Println("result is division -> ", division)
	} else {
		fmt.Println("result is division ->", division,
			"remainder ->", remainder)
	}
}

func demo(name string) int {
	fmt.Println(name)
	return 1 + 2
}

func intDivision(num1 int, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	division := num1 / num2
	remainder := num1 % num2
	return division, remainder, err
}
