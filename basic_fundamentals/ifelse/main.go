package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome! to ifelse! world!")
	reader := bufio.NewReader(os.Stdin)
	var input string
	var count int = 0
	for {
		input, _ = reader.ReadString('\n')
		fmt.Println(input)
		count++
		if count > 3 {
			break
		}
	}
}
