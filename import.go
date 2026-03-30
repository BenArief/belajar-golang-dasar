package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main(){
	result := helper.SayHello("Ben")
	fmt.Println(result)

	fmt.Println(helper.Application)
	helper.Contoh()
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye("Ben"))
}