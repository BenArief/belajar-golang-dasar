package main

import "fmt"

func getFullName() (string, string) {
	return "Ben", "Arief"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}