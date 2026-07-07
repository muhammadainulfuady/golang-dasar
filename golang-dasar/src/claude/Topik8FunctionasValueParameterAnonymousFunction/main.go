package main

import "fmt"

func main() {
	// TUGAS 1: buat 2 anonymous function, simpen ke variabel:
	// - kalikan: nerima 2 int, return hasil perkalian
	// - kurangi: nerima 2 int, return hasil pengurangan
	kalikan := func(a, b int) int {
		return a * b
	}

	kurangi := func(a, b int) int {
		return a - b
	}

	// TUGAS 2: panggil prosesAngka() 2 kali, sekali pakai "kalikan", sekali pakai "kurangi"
	// print hasil dari masing-masing pemanggilan

	perkalian := prosesAngka(10, 5, kalikan)
	pengurangan := prosesAngka(10, 5, kurangi)
	fmt.Printf("Hasil dari perkalian : %d\n", perkalian)
	fmt.Printf("Hasil dari pengurangan : %d\n", pengurangan)

	// TUGAS 3 (bonus): buat 1 anonymous function LANGSUNG dipanggil di tempat
	// (tanpa disimpen ke variabel dulu), yang ngeprint "Halo dari anonymous function!"
	func() {
		fmt.Println("Halo dari anonymous function!")
	}()
}

// function ini nerima function lain sebagai parameter
func prosesAngka(a, b int, operasi func(int, int) int) int {
	return operasi(a, b)
}
