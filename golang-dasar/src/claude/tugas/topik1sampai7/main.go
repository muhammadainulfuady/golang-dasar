package main

import "fmt"

const BATAS_MINIM_STOK = 10

func main() {
	// data penjualan: nama produk -> jumlah terjual
	penjualan := map[string]int{
		"Kopi":  45,
		"Gula":  8,
		"Teh":   30,
		"Susu":  5,
	}

	// TUGAS 1: 
	// Loop semua penjualan pakai for+range
	// Kalau jumlah terjual < BATAS_MINIM_STOK, print "<nama>: perlu restock segera"
	// Selain itu, print "<nama>: stok masih aman"

	// TUGAS 2:
	// Panggil function totalDanTertinggi() di bawah, kasih argumen dari SEMUA
	// nilai yang ada di map penjualan (pakai variadic + slice, boleh kumpulin
	// nilainya ke slice dulu sebelum dipanggil)
	// Print hasil total dan tertingginya

	// TUGAS 3:
	// Pakai switch, buat kategori toko berdasarkan total penjualan:
	// >= 100 -> "Toko Ramai"
	// >= 50  -> "Toko Sedang"
	// selain itu -> "Toko Sepi"
	// Print hasil kategorinya
}

// TUGAS 4: lengkapi function ini
// - Terima variadic int
// - Return DUA nilai: total semua angka, dan nilai TERTINGGI
func totalDanTertinggi(angka ...int) (total int, tertinggi int) {
	// isi logic di sini
}