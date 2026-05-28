package main

import "fmt"

func main() {
	siswas := map[string]int{
		"Budi":     80,
		"Andi":     70,
		"Siti":     85,
		"Ainul":    90,
		"Fuady":    88,
		"Muhammad": 56,
	}

	totalNilai := 0

	maks := 0
	minimal := 100

	fmt.Println("Daftar siswa:")

	for nama, nilai := range siswas {

		fmt.Println(nama, ":", nilai)

		totalNilai += nilai

		if nilai > maks {
			maks = nilai
		}

		if nilai < minimal {
			minimal = nilai
		}
	}

	fmt.Println()

	for nama, nilai := range siswas {
		if nilai >= 75 {
			fmt.Println(nama, ": Lulus")
		} else {
			fmt.Println(nama, ": Tidak lulus")
		}
	}

	fmt.Println()
	fmt.Println("Total nilai :", totalNilai)
	fmt.Println("Rata rata :", totalNilai/len(siswas))
	fmt.Println("Nilai tertinggi :", maks)
	fmt.Println("Nilai terendah :", minimal)
}