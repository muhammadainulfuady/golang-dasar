package main

import "fmt"

func main() {
	fmt.Println(pangkat(2, 5)) // 2 pangkat 5 = 32
	fmt.Println(pangkat(3, 3)) // 3 pangkat 3 = 27
	fmt.Println(pangkat(5, 0)) // apapun pangkat 0 = 1
}

// TUGAS: lengkapi function recursive ini
// pangkat(basis, eksponen) menghitung basis^eksponen
// contoh: pangkat(2, 3) = 2 * 2 * 2 = 8
//
// Petunjuk base case: kalau eksponen == 0, hasilnya selalu 1 (aturan matematika)
// Petunjuk recursive case: basis * pangkat(basis, eksponen-1)
func pangkat(basis, eksponen int) int {
	if eksponen == 0 {
		return 1
	}
	return basis * pangkat(basis, eksponen-1)
}
