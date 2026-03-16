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

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	var daySlice1 = days[5:] // Sabtu dan minggu
	fmt.Println(daySlice1)
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days) // data di array juga berubah karena slice mereferensikan data di array

	daySlice2 := append(daySlice1, "Libur baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2) // data di slice 1 tidak berubah karena append membuat array baru
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "ilham"
	newSlice[1] = "ilham"
	// newSlice[2] = "ilham"  error karena kapasitas slice hanya 5, index 2 tidak bisa diisi
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "ilham2")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}