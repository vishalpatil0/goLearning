package utility

import "fmt"

func Display() {
	fmt.Println("This is a utility function")
}

func Add5(i *int) {
	*i = *i + 5
}
