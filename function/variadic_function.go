package main

import (
	"fmt"

)

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func sumAllSlice(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	sum := sumAll(10, 10, 10, 10, 10)
	fmt.Println(sum)

	fmt.Println("Ini hasil Slice : ", sumAllSlice([]int{1, 1, 1, 1, 1}))

	numbers2 := []int{10,10,10}
	fmt.Println(sumAll(numbers2...))
}