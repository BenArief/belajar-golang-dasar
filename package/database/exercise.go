package main

import "fmt"

func main() {
	//Variables
	// var myNum int = 90
	// var myWord string = "hello World"
	// var myBool bool = true
	// x := 12.75

	// fmt.Println(myNum)
	// fmt.Println(myWord)
	// fmt.Println(myBool)
	// fmt.Printf("x: %.2f\n", x )

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

	//slice
	slice1 := []int{1,2,3} // kantong elastis
	fmt.Println(len(slice1)) // isi nya berapa
	fmt.Println(cap(slice1)) // kapasitas nya berapa, bisa lebih besar dari panjang nya

	//bikin array dari slice
	ary1 := [8]int{1,2,3,4,5,6,7,8}
	slice_from_array := ary1[2:4]

	fmt.Println("slice_from_array:", slice_from_array)
	fmt.Println("Length: ", len(slice_from_array))
	fmt.Println("Capacity: ", cap(slice_from_array))

	//using make slice
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

}