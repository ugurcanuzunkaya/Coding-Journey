package main

import (
	"fmt"
)

func main() {
	var name string
	// Variable name is declared as a string

	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	// Output: Hello, name

	var age int
	// Variable age is declared as an integer

	fmt.Println("What is your age?")
	fmt.Scan(&age)
	fmt.Printf("%s's age is %d \n", name, age)
	// Output: name's age is age

	newMessage := fmt.Sprintf("Welcome to your %d age %s!", age, name)
	fmt.Println(newMessage)
	// Output: Welcome to your age name!
}
