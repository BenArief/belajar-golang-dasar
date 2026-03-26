package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Error bro")
	}

	// jangan taruh recover setelah panic karna ga akan dieksekusi
	// message := recover()
	// println("Terjadi Panic", message)
}

func main(){
	runApp(true)
}

