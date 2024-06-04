package main

import (
	"fmt"
)

func main() {
	// printing
	fmt.Println("hellow world")

	// variable
	var a int = 12
	var b int = 13
	fmt.Println(a + b)

	// int8 int16 int32 int64
	var c int8 = 0 + 127
	fmt.Println(c)

	// unsinged int
	var d uint64 = 9999999999999999999
	var e int64 = 9199999999999999999
	// fmt.Println(d+e) mismatched type
	// solution: type cast
	println(d, e)
	var f float32 = 34.3
	var g int16 = int16(f)
	println(g)

	//string

	var str string = "hello world"
	println(str)
	str = "hello \n world"
	println(str)
	str = `hello 
	world`
	println(str)

	//default value for int, float and rune is 0
	// for string "" emtpry string
	//for boolean false

	//type inferencing
	var name = "vishal"
	println(name)

	myVar := 12
	println(myVar)

	var z, x, w = 1, "vishal", 45.4

	println(z, x, w)

	const myConst string = "name"

	FunctionInFile1()

}
