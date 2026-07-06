package main

import "fmt"

func main() {
	umur := 20
	punyaSIM := true
	saldoDompet := 50000
	hargaTiket := 75000

	// TUGAS:
	// 1. Cek apakah boleh nyetir sendiri: umur harus >= 17 DAN punyaSIM harus true
	// 2. Cek apakah cukup uang buat beli tiket: saldoDompet harus >= hargaTiket
	// 3. Cek apakah TIDAK cukup uang (pakai operator ! dari hasil poin 2)
	// 4. Print ketiga hasil boolean itu

	// Bonus: cari sisa bagi (modulus) dari saldoDompet dibagi hargaTiket, print hasilnya
	boolNyetir := umur >= 17 && punyaSIM
	boolTiket := saldoDompet >= hargaTiket
	boolTidakCukup := !boolTiket
	modSaldoDompet := saldoDompet % hargaTiket

	fmt.Printf("Apakah saya boeleh nyetir sendiri : %t\n", boolNyetir)
	fmt.Printf("Apakah saldo cukup untuk membeli tiket : %t\n", boolTiket)
	fmt.Printf("Saya memang tidak bisa beli tiket : %t\n", boolTidakCukup)
	fmt.Printf("Sisa bagi dari saldo dompet : %d\n", modSaldoDompet)
}
