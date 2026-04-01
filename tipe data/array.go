package main

import "fmt"

func main2() {
	var names [3]string
	names[0] = "Windah"
	names[1] = "Charles"
	names[2] = "Jeremy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		100,
		200,
		65,
		12,
	}
	fmt.Println(values) 
	fmt.Println(len(values))
	
}