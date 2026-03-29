package main

import (
	"fmt"
)

func random() any {
	return true
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// var hai string 
	// fmt.Println(hai)	

	// Jangan konversiin ke tipe data yang lain apabila lu ga tau tipe data yang sebenarnya apa
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// instead of ga ada kepastian, kita pake switch aja
	switch value := result.(type) {
		case string :
			fmt.Println("String :", value)
		case int :
			fmt.Println("Int :", value)
		default :
		fmt.Println("Unknown Type :", value)

	}
}