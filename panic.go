package main

import (
	"fmt"
	"internal/poll"

	"golang.org/x/tools/go/analysis/passes/modernize"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

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

