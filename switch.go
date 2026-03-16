package main

import "fmt"

func main() {
	name := "Windah"

	switch name {
	case "Ben":
		fmt.Println("Hello Ben")
	case "Windah":
		fmt.Println("Hai Windah")
	default:
		fmt.Println("Hello Stranger")
	}
}