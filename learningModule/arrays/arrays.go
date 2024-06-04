package main

import "fmt"

func main() {
	var arr [3]int64

	fmt.Println("0th index", arr[0])

	arr[2] = 123

	fmt.Println(arr[2])

	fmt.Println(arr[0:3])

	//printing memory location
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	//direct initilization
	var arr1 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(arr1[1])

	//slices is just a array with empty size declaration.
	var arraySlices []int32 = []int32{1, 2, 3}
	arraySlices = append(arraySlices, 7)
	fmt.Println(arraySlices)
	fmt.Println("length of array -> ", len(arraySlices), "\n size of array -> ", cap(arraySlices))

	//map
	var map1 map[string]uint16 = make(map[string]uint16)
	map1["key"] = 12
	fmt.Println(map1["key"])

	var map2 = map[string]uint16{"name": 1, "s": 2}

	for k, v := range map2 {
		fmt.Printf("key : %v and value : %v", k, v)
	}
	println()
	var val, present = map2["s"]
	if present {
		fmt.Println("value is present -> ", val)
	} else {
		fmt.Println("value is absent")
	}
	delete(map2, "s")
	val, present = map2["s"]
	if present {
		fmt.Println("value is present -> ", val)
	} else {
		fmt.Println("value is absent")
	}

	//for loop

	println()
	println()
	println()
	println("for loops ------>")

	i := 0

	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	i = 5
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	age := &i

	fmt.Println("age -> ", *age)
}
