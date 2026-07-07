package main

import "fmt"

func main() {
	untung, rugi := hitungLabaRugi(500000, 350000)
	fmt.Println("Untung:", untung)
	fmt.Println("Rugi:", rugi)
}

// TUGAS: lengkapi function ini
// - Kalau modal (parameter ke-2) lebih kecil dari pendapatan (parameter ke-1) -> untung = pendapatan - modal, rugi = 0
// - Kalau modal lebih besar dari pendapatan -> rugi = modal - pendapatan, untung = 0
// - Kalau sama -> untung = 0, rugi = 0
// - WAJIB pakai NAMED RETURN VALUE (bukan return biasa)
func hitungLabaRugi(pendapatan int, modal int) (untung int, rugi int) {
	// isi logic di sini
	if modal < pendapatan {
		untung = pendapatan - modal
		rugi = 0
	} else if modal > pendapatan {
		rugi = modal - pendapatan
		untung = 0
	} else {
		untung = 0
		rugi = 0
	}
	return
}
