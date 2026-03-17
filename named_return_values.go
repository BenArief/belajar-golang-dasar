package main

import "fmt"

func getCompleteName() (firstName string, lastName string) {
	firstName = "Ben"
	lastName = "Arief"

	return firstName, lastName
}

func main() {
	a, b := getCompleteName()
	fmt.Println(a, b)
}