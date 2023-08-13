package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Enter task: // List Task: // Quit

	// First thing is to welcome
	fmt.Println("Hello! Welcome to Task-Manager {Terminal Mode}")

	tasks := make([]string, 0) // to store the tasks in it with the string format starting with value 0 or index 0

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n1. Enter Task\n2. List Task\n3. Quit")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Enter Task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)
			// fmt.Println("Task Added!")
		case "2":
			fmt.Print("\n These are the following tasks")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		default:
			fmt.Print("Incorrect value!!")
		}
	}
}
