package main

import "fmt"

func main() {
	fmt.Println("Welcome to the basic fundamentals")

	var arr [2]string
	arr[0] = "Tejas"
	arr[1] = "Kulkarni"
	// in go there are no while loops and for each etc like all other major languages
	// 1. for{} this loop is infinite
	// 2. for p1, p2:= range declared variable{}
	// 3. for int i:=0 i<n i++{}

	// lets use the range loop
	for _, value := range arr {
		fmt.Println(value)
	}
}
