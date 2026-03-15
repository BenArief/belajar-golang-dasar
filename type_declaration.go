package main

import "fmt"

func main() {
	type noKTP string 

	var ktpBen noKTP = "1234232"
	var contoh string  = "00000"
	var contohKtp noKTP = noKTP(contoh)

	fmt.Println(ktpBen)
	fmt.Println(contohKtp)
}