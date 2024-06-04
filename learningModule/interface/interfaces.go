package main

type Animal interface {
	eat()
}

// embedded interface, interface extending another interface
type Bird interface {
	fly()
	Animal
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) eat() {
	println("Dog is eating")
}

func feed(a Animal) {
	a.eat()
}

func capture(b Bird) {
	b.eat()
	b.fly()
}

type Parrot struct {
	name string
}

func (p Parrot) fly() {
	println("Parrot is flying")
}

func (p Parrot) eat() {
	println("Parrot is eating")
}

func check(value interface{}) {
	switch value.(type) {
	case int:
		println("Integer")
	case string:
		println("String")
	default:
		println("Unknown")
	}
}

func validate(value any) {
	i, isInt := value.(int)
	if !isInt {
		println("Invalid")
		return
	} else {
		println(i)
	}
}
func main() {
	var d Dog = Dog{name: "Buddy"}
	feed(d)
	// var c Cat = Cat{name: "Kitty"}
	// feed(c)
	// struct need to implement method of interface

	check(10)
	check("Hello")
	check(10.5)

	var p Parrot = Parrot{name: "Polly"}
	capture(p)

	validate(10.4)
}
