package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++{
		fmt.Println("Perluangan ke", counter)
	}

	fmt.Println("Selesai!")

	names := []string{"Ben", "Arief", "Moggalana"}
	for index, name := range names{
		fmt.Println("Index ke-", index+1, "=", name)
	}
}