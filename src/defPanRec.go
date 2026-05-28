package main

import "fmt"

func main() {
	hasil := runApp(10, 0)
	fmt.Println("Hasil:", hasil)
}

func endApp() {
	fmt.Println("Function selesai")
	messageErr := recover()
	if messageErr != nil {
		fmt.Println("Terjadi eror : ", messageErr)
	}
}

func runApp(a int, b int) int {
	defer endApp()
	if b == 0 || a == 0 {
		panic("tidak bisa membagi 0")
	}
	return a / b
}
