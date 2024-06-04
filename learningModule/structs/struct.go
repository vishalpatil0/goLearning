package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type student struct {
	Name string
	Age  int
}

type monitor struct {
	names string
	student
}

func (std *student) printData() {
	fmt.Println(std.Name)
	fmt.Println(std.Age)
}

func (std *student) nullifiedData() {
	std.Name = ""
	std.Age = 0
}

func newStudent(name string, age int) (*student, error) {
	if name == "" && age <= 0 {
		return nil, errors.New("Invalid data")
	}
	return &student{
		Name: name,
		Age:  age,
	}, nil
}

func newMonitor() monitor {
	return monitor{
		names: "John",
		student: student{
			Name: "John",
			Age:  20,
		},
	}
}

func main() {
	var std student = student{Name: "John", Age: 20}
	std.printData()
	std.nullifiedData()
	std.printData()

	std1, err := newStudent("", -23)
	if err != nil {
		fmt.Println(err)
	} else {
		std1.printData()
	}

	mon := newMonitor()
	mon.printData()

	//saving in json file
	jsn, err := json.Marshal(std)
	if err != nil {
		fmt.Println(err)
	} else {
		os.WriteFile("js.json", jsn, 0644)
	}
}
