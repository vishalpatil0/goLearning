package main

import "fmt"

type struct1 struct {
	name string
	age  uint8
}

// structure method
func (e struct1) display(name string) {
	fmt.Println(name)
}

func main() {
	str11 := struct1{"s", 1}
	fmt.Println(str11.age)

	//anonymous struct
	var str32 = struct {
		name string
	}{"name"}
	fmt.Println(str32)
	str11.display("name")
}
