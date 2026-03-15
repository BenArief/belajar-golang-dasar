package main

import "fmt"

func main() {
	const (
		firstName = "Ben"
		lastName  = "Arief"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Error
	// fmt.Println(fistName)
	// fistName = "ilham"
	// fmt.Println(fistName)
}
