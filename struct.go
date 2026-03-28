package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello,", name , "my name is ", customer.Name)
}

func (c Customer) tellAddress(){
	fmt.Println("I'm live in", c.Address)
}

func main() {
	var ben Customer
	fmt.Println(ben)

	ben.Name = "Ben Arief"
	ben.Address = "Tangerang"
	ben.Age = 21

	fmt.Println(ben)
	fmt.Println(ben.Address)

	joko := Customer{
		Name: "Joko Gunawan",
		Address: "Indonesia",
		Age: 30,
	}

	fmt.Println(joko)

	// struct literal
	sarah := Customer{"Sarah", "Indonesia", 25}
	fmt.Println(sarah)	

	sarah.sayHello("Agus")
	joko.sayHello("Budi")

	ilham := Customer{
		Name: "Ilham Kurniawan",
		Address: "Bekasi",
		Age: 28,
	}

	ilham.sayHello("Windah")
	ilham.tellAddress()
}