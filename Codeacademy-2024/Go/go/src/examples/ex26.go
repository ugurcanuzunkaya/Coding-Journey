package main

import "fmt"

func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"

	// if lockCombo is equal to robAttempt, print "The vault is now opened."
	if lockCombo == robAttempt {
		fmt.Println("The vault is now opened.")
	} else {
		// if lockCombo is not equal to robAttempt, print "Police are informed!"
		fmt.Println("Police are informed!")
	}
}
