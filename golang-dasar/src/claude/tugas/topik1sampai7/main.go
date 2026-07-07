package main

import (
	"fmt"
)

const BATAS_MINIM_STOK = 10

func main() {
	// data penjualan: nama produk -> jumlah terjual
	penjualan := map[string]int{
		"Kopi": 45,
		"Gula": 8,
		"Teh":  30,
		"Susu": 5,
	}

	// TUGAS 1:
	// Loop semua penjualan pakai for+range
	// Kalau jumlah terjual < BATAS_MINIM_STOK, print "<nama>: perlu restock segera"
	// Selain itu, print "<nama>: stok masih aman"
	for key, value := range penjualan {
		if value < BATAS_MINIM_STOK {
			fmt.Printf("%s : perlu restock segera\n", key)
		} else {
			fmt.Printf("%s : stok masih aman\n", key)
		}
	}

	// TUGAS 2:
	// Panggil function totalDanTertinggi() di bawah, kasih argumen dari SEMUA
	// nilai yang ada di map penjualan (pakai variadic + slice, boleh kumpulin
	// nilainya ke slice dulu sebelum dipanggil)
	// Print hasil total dan tertingginya

	var temp []int
	for _, nilai := range penjualan {
		temp = append(temp, nilai)
	}
	total, tertinggi := totalDanTertinggi(temp...)
	fmt.Printf("\nTotal terjual : %d\n", total)
	fmt.Printf("Penjualan tertinggi : %d\n", tertinggi)

	// TUGAS 3:
	// Pakai switch, buat kategori toko berdasarkan total penjualan:
	// >= 100 -> "Toko Ramai"
	// >= 50  -> "Toko Sedang"
	// selain itu -> "Toko Sepi"
	// Print hasil kategorinya

	switch {
	case total >= 100:
		fmt.Println("\nToko ramai")
	case total >= 50:
		fmt.Println("\nToko sedang")
	default:
		fmt.Println("\nToko sepi")
	}
}

// TUGAS 4: lengkapi function ini
// - Terima variadic int
// - Return DUA nilai: total semua angka, dan nilai TERTINGGI
func totalDanTertinggi(angka ...int) (total int, tertinggi int) {
	if len(angka) == 0 {
		return 0, 0
	}
	tertinggi = angka[0]
	for _, nilai := range angka {
		total += nilai
		if nilai > tertinggi {
			tertinggi = nilai
		}
	}
	return
}
