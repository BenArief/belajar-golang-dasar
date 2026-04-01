package main

import "fmt"

type Blaklist func(string) bool

func registerUser(name string, blacklist Blaklist) {
	if blacklist(name) {
		fmt.Println("You're Blocked!", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func (name string) bool  {
		return name == "anjing"
	}
	registerUser("Ben", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}