package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // Hanya copy value, bukan reference

	// address2.City = "Bandung"
	// fmt.Println(address1) // ga berubah
	// fmt.Println(address2) // subang jadi bandung

	// Kalau mau persis sama, gunakan &
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1) // bandung
	fmt.Println(address2) // bandung juga
}