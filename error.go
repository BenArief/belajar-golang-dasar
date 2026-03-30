package main

import (
	"errors"
	"fmt"
)

// type error interface {
// 	Error() string
// }

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol tidak bisa dioperasikan!")
	} else {
		return  nilai / pembagi, nil
	}
}

func main(){
	hasil, err := Pembagian(4,0)
	if err == nil {
		fmt.Println("Hasil =", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}