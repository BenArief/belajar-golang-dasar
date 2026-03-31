package main

import "fmt"

func main() {
	//Variables
	var myNum int = 90
	var myWord string = "hello World"
	var myBool bool = true
	x := 12.75

	fmt.Println(myNum)
	fmt.Println(myWord)
	fmt.Println(myBool)
	fmt.Printf("x: %.2f\n", x )

	//Array
	selamat := [4]string{"Pagi", "Siang", "Sore", "Malam"}
	arr2 := [...]int{1,2,3,4,5,2,1,2,3,3,3,4,5,6,7,8,9,0}
	fmt.Println(selamat)
	fmt.Println(arr2)

	var cars = [4]string{"volvo", "Mercedes", "BYD", "Tesla"}
	fmt.Println(cars)
	fmt.Println(cars[2])
	//change value
	cars[0] = "Rolls-Royce"
	fmt.Println(cars)

	arr3 := [5]int{1:5, 3:100}
	fmt.Println(arr3)

	//length of array
	fmt.Println("Length of cars:", len(cars))
	fmt.Println("Length of arr2:", len(arr2))
}