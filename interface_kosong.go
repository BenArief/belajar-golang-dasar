package main

import "fmt"

func ups() any {
	// return 1
	// return true
	return "Ups"
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
}