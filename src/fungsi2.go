package main

import "fmt"

func main() {
	result := angka(12, checkAngka)
	if result == 0 {
		fmt.Println("genap bos")
	} else {
		fmt.Println("ganjel bos")
	}

	jancik := func(ancik string) {}
}

type N func(int) int

func angka(number int, n N) int {
	check := n(number)
	return check

}

func checkAngka(n int) int {
	if n%2 == 0 {
		return 0
	} else {
		return 1
	}
}
