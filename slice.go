package main

import (
	"fmt"
)

func main() {
	var months = [...]string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}

	slice1 := months[4:7]
	fmt.Println(slice1) // data dari array index 4 sampe 6, index 7 tidak termasuk

	slice2 := months[:3]
	fmt.Println(slice2) // data dari array index 0 sampe 2

	slice3 := months[5:]
	fmt.Println(slice3) // data dari array index 5 sampe akhir

	slice4 := months[:]
	fmt.Println(slice4) // semua data di array
}