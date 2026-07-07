package main

import "fmt"

func main() {
	// TUGAS 1: panggil cariNilaiTertinggi dengan beberapa angka langsung
	tertinggi := cariNilaiTertinggi(80, 95, 70, 88)
	fmt.Println("Nilai tertinggi:", tertinggi)

	// TUGAS 2: panggil cariNilaiTertinggi lagi, tapi dari SLICE yang udah ada
	// (pakai teknik "membongkar" slice dengan ...)
	nilaiSiswa := []int{60, 75, 92, 81}
	tertinggi2 := cariNilaiTertinggi(nilaiSiswa...)
	fmt.Println("Nilai tertinggi dari slice:", tertinggi2)
}

// TUGAS 3: lengkapi function ini
// - Terima jumlah angka berapapun (variadic)
// - Cari dan return nilai TERTINGGI dari semua angka yang dikasih
// - Kalau gak ada angka sama sekali (variadic kosong), return 0
func cariNilaiTertinggi(angka ...int) int {
	// isi logic di sini
	if len(angka) == 0 {
		return 0
	}
	tinggi := angka[0]
	for _, nilai := range angka {
		if nilai > tinggi {
			tinggi = nilai
		}
	}
	return tinggi
}
