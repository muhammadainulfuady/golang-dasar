package main

import "fmt"

func main() {
	// TUGAS:
	// 1. Buat constant TARIF_PAJAK = 0.11 (11%, PPN)
	// 2. Buat variable hargaBarang (terserah tipe & cara declare-nya) = 200000
	// 3. Hitung pajak = hargaBarang * TARIF_PAJAK
	// 4. Hitung totalBayar = hargaBarang + pajak
	// 5. Print: "Harga: Rp <harga>, Pajak: Rp <pajak>, Total: Rp <total>"

	// Bonus (opsional): coba sengaja ubah nilai TARIF_PAJAK setelah dideklarasikan,
	// lihat apa yang terjadi, terus hapus lagi baris itu setelah kamu liat errornya

	const TARIF_PAJAK = 0.11
	hargaBarang := 200000
	pajak := float64(hargaBarang) * TARIF_PAJAK
	totalBayar := float64(hargaBarang) + pajak
	fmt.Printf("Harga: Rp %d, Pajak: Rp %.2f, Total: Rp %.2f", hargaBarang, pajak, totalBayar)
}
