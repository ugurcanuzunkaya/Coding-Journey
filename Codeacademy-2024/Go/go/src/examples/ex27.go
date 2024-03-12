package main

import "fmt"

func main() {
	vaultAmt := 2356468

	// vaultAmt is greater than 200000 prints "We're going to need more bags."
	if vaultAmt >= 200000 {
		fmt.Println("We're going to need more bags.")
	}
}
