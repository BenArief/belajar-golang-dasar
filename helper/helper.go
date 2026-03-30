package helper

import "fmt"

//huruf kecil tidak bisa diakses di luar package
var version = "1.0.0"
var Application = "Golang"

func sayGoodBye (name string) string{
	return "Good Bye " + name
}

// huruf besar bisa diakses di luar package
func SayHello(name string) string{
	return  "Hello " + name
}

func Contoh() {
	fmt.Println(version)
	fmt.Println(sayGoodBye("Ben"))
}