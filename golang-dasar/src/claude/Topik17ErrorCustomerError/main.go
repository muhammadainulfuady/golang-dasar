package main

import (
	"errors"
	"fmt"
)

var (
	ErrPembagian = errors.New("Tidak bisa di bagi dengan nol")
)

func bagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, ErrPembagian
	}
	return nilai / pembagi, nil
}

func main() {
	if result, err := bagi(10, 0); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Hasil dari pembagian adalah :", result)
	}
}
