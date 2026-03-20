package main

import (
	"fmt"
	"strings"
)

func test1() {
	var halo string = "Ini angka 2"
	fmt.Println(halo)
	arrai := [5]int{1,5,3,6}
	for i := 0; i < len(arrai); i++ {
		fmt.Println("Value", i, ":", arrai[i])
	}
}

func main(){
	test1()
	num := 3.15
	fmt.Printf("%f", num)
}