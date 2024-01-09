package main

import "fmt"

func main(){
	
	// Single Input Code
	fmt.Println("Enter name:")
	
	var name string;
	fmt.Scanf("%s", &name);

	fmt.Println("Hello! ", name);
	// The above scanf will only take a single input without space
	// If there are two inputs having the space then it will only print the first word/letter
	

	// #Multiple Inputs
	var gender string
	var is_male bool

	fmt.Println("Enter your name and are you male?")
	fmt.Scanf("%s %t", &gender,&is_male);

	fmt.Println("Your output is: ", gender, is_male);

	
}
