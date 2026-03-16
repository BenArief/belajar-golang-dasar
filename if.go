package main

import "fmt"

func main() {
	name := "Ben"

	if name == "Ben" {
		fmt.Println("Hello Ben")
	} else if name == "Windah" {
		fmt.Println("Hai Windah")
	} else {
		fmt.Println("Hello Stranger")
	}

	if length := len(name); length > 5{
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
