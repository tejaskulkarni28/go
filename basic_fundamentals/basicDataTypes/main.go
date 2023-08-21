package main

import "fmt"

func main() {
	// type 1
	var a = "a"
	fmt.Println(a)

	// type 2
	var b string
	b = "b"
	fmt.Println(b)

	// type 3
	var c = 101
	fmt.Println(c)

	// type 4
	var d int
	d = 102
	fmt.Println(d)

	// type 5
	const e = 103
	// e = 102 cannot assign because of declaration of const for e variable
	fmt.Println(e)

	// type 6
	f := "f" // this is auto declaration also called syntactic sugar
	fmt.Println(f)

	// slices and arrays
	// arrays have the defined size
	// slices have but are re-sizable yooooo

	// declaring the array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("Looping using loop range")
	for index, value := range arr {
		fmt.Println(index, " ", value)
	}

	fmt.Println("Looping using basic loop style")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("Looping using infinite loop style")
	for {
		fmt.Println("Hello! This is a infinite loooooppppppp!!!!!!!!!!!!!!!!")
		break // used break because will be piss my system xD
	}
}
