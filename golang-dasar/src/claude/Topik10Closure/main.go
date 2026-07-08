package main

import "fmt"

// buatTabungan menerima saldo awal, dan mereturn sebuah closure function
func buatTabungan(saldoAwal int) func(int) int {
	// TUGAS 1: Return sebuah anonymous function yang menerima 1 parameter int (misal: 'setoran').
	// Di dalam function tersebut, tambahkan 'saldoAwal' dengan 'setoran'.
	// Lalu, return nilai 'saldoAwal' yang terbaru.
	return func(setoran int) int {
		saldoAwal += setoran
		return saldoAwal
	}
}

func main() {
	// TUGAS 2: Panggil buatTabungan dengan saldo awal 10000, lalu simpan return value-nya
	// ke dalam variabel bernama 'tabunganku'.
	tabunganku := buatTabungan(10000)

	// TUGAS 3: Lakukan 2 kali setoran dengan memanggil 'tabunganku':
	// - Setoran pertama: 5000. Print hasilnya (harus 15000).
	// - Setoran kedua: 2000. Print hasilnya (harus 17000).
	setoranPertama := tabunganku(5000)
	fmt.Println(setoranPertama)
	setoranKedua := tabunganku(2000)
	fmt.Println(setoranKedua)
}
