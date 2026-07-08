package main

import "fmt"

func main() {
	fmt.Println("=== Percobaan 1 ===")
	hasil1 := bagiAman(10, 2)
	fmt.Println("Hasil:", hasil1)

	fmt.Println("\n=== Percobaan 2 ===")
	hasil2 := bagiAman(10, 0)
	fmt.Println("Hasil:", hasil2)

}

// TUGAS: lengkapi function ini
// - Kalau b == 0, ini bakal nge-trigger panic bawaan Go (divide by zero) waktu dieksekusi
// - Pakai defer + recover buat "menangkap" panic itu, biar program GAK CRASH
// - Kalau ke-recover, print pesan "Error: pembagian dengan nol dicegah" dan return 0
// - Kalau normal (gak panic), return hasil pembagian biasa
func bagiAman(a, b int) (hasil int) {
	// isi logic di sini
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Eror: pembagian dengan nol dicegah")
			hasil = 0
		}
	}()

	hasil = a / b
	return
}
