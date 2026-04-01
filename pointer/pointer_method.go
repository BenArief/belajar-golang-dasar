package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	ben := Man{"Ben"}
	ben.Married()
	fmt.Println(ben.Name)
}