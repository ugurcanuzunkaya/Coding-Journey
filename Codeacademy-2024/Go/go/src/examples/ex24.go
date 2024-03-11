package main

import "fmt"

func main() {
	// heistReady is a boolean variable
	heistReady := true

	// if heistReady is true, print "Let's Go!"
	if heistReady {
		fmt.Println("Let's Go!")
	}
	// Output: Let's Go!

	// if heistReady is false, print "Act normal"
	heistReady = false

	if heistReady {
		fmt.Println("Act normal")
	}
	// Output: Act normal

}
