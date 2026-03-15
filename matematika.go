package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 23
	var c = a * b % d
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println("i = ", i)

	var j = 1
	j++
	j++
	fmt.Println("j = ", j)
}