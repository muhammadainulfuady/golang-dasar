package main

import "fmt"

func main() {
	// TUGAS:
	// 1. Buat SLICE bernama nilaiUjian, isi awal: 80, 75, 90
	// 2. Tambahin 1 nilai lagi pakai append(): 85
	// 3. Buat MAP bernama biodata, isinya:
	//    "nama" -> "Andi"
	//    "kelas" -> "12 IPA"
	// 4. Print semua isi nilaiUjian
	// 5. Print biodata["nama"] dan biodata["kelas"]
	// 6. Hitung dan print PANJANG (jumlah elemen) dari nilaiUjian pakai len()

	nilaiUjian := []int{
		80, 75, 90,
	}
	nilaiUjian = append(nilaiUjian, 85)

	fmt.Println("Nilai ujian : ", nilaiUjian)
	
	biodata := map[string]string{
		"nama":  "Andi",
		"kelas": "12 IPA",
	}
	fmt.Printf("Nama : %s\nKelas : %s\n", biodata["nama"], biodata["kelas"])
	fmt.Println("Panjang dari slice nilai ujian adalah", len(nilaiUjian))
}
