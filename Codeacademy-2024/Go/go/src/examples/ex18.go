package main

import "fmt"

func main() {
	fmt.Println("Let's first see how", "the Println() method works.")
	fmt.Println("Notice that each statement adds a newline for us.")
	fmt.Println("There's also a default space", "between the string arguments.")
	// Output: Let's first see how the Println() method works.
	// Notice that each statement adds a newline for us.
	// There's also a default space between the string arguments.

	fmt.Println("Println", "adds", "spaces and newline")
	// Output: Println adds spaces and newline

	// Add your code below:
	fmt.Print("Print", "does", "not", "add", "spaces")
	fmt.Print("Or newlines unless specified \n")
	// Output: PrintdoesnotaddspacesOr newlines unless specified
}
