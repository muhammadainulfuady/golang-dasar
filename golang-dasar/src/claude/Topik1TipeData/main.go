package main

import "fmt"

func main() {
	var namaBarang string = "Kopi"
	var hargaSatuan int = 15000
	var jumlahBeli float64 = 2.5 // anggap ini "2.5 kg" kopi
	var adaDiskon bool = true

	// TUGAS:
	// 1. Hitung total harga = hargaSatuan * jumlahBeli
	//    (perhatiin: ini bakal error kalau langsung dikaliin, kenapa? benerin caranya)
	// 2. Kalau adaDiskon == true, total harga dipotong 10%
	// 3. Print hasil akhir dengan format: "Total belanja Kopi: Rp <total>"
	totalHarga := float64(hargaSatuan) * jumlahBeli
	diskon := (totalHarga * 10) / 100
	if adaDiskon {
		fmt.Printf("Total harga belanja %s : Rp %.2f", namaBarang, totalHarga-diskon)
	} else {
		fmt.Printf("Total harga belanja %s : Rp %.2f", namaBarang, totalHarga)
	}
}
