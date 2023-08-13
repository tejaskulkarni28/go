package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome! to Task List Manager")
	// here the tasks is a declaration of the variable and make is function which creates a slice
	// and slice in golang is an array like a dynamic structure which is re-sizable
	// and the first args is the type and the second args is the intial value.
	tasks := make([]string, 0)

	// here the reader is the variable and the bufio is the buffered reader which reads the os.standard input from the user and holds in the reader variable
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n1. Add Task\n2. List Task\n3. Quit\nChoose an option: ")
		// here ,_ is user for to avoid the \n for avoiding the error or unwanted elements
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)
		case "2":
			fmt.Println("\nTasks: ")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case "3":
			fmt.Println("GoodBye!! This was from golang :)")
			return
		default:
			fmt.Print("Invalid Option: Please Choose Again")
		}
	}

}
