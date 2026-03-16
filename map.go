package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{
	// 	"name": "Ben"
	// 	"address": "Tangerang"
	// }

	person := map[string]string{
		"name":    "Ben",
		"address": "Tangerang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := map[string]string{
		"title": "Buku Golang",
		"author": "Windah Basudara",
		"gatau": "gatau2",
	}
	fmt.Println(book)
	delete(book, "gatau")
	fmt.Println(book)
}