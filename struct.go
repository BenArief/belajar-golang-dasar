package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ben Customer
	fmt.Println(ben)

	ben.Name = "Ben Arief"
	ben.Address = "Tangerang"
	ben.Age = 21

	fmt.Println(ben)
	fmt.Println(ben.Address)

	joko := Customer{
		Name: "Joko Gunawan",
		Address: "Indonesia",
		Age: 30,
	}

	fmt.Println(joko)

	// struct literal
	sarah := Customer{"Sarah", "Indonesia", 25}
	fmt.Println(sarah)	
}