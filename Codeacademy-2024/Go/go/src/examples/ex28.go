package main

import "fmt"

func main() {
	rightTime := true
	rightPlace := true

	// If it's the right time and the right place, we're outta here!
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}

	enoughRobbers := false
	enoughBags := true

	// If there are enough robbers and enough bags, grab everything
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
}
