package main

import "fmt"

func main() {
	siswas := map[string][]int{
		"Budi":     {80, 90, 70},
		"Andi":     {60, 75, 80},
		"Siti":     {95, 88, 92},
		"Ainul":    {70, 72, 68},
		"Muhammad": {85, 78, 90},
	}
	showSiswa(siswas)
	fmt.Println("")
	showNilai(siswas)
}

func showSiswa(siswa map[string][]int) {
	for key, value := range siswa {
		fmt.Println(key, ":", value)
	}
}

func showNilai(siswa map[string][]int) {
	for key, value := range siswa {
		total, avg := sumAndAvgV2(value...)
		fmt.Println(key, ":")
		fmt.Println("Nilai :", total)
		fmt.Println("Avg : ", avg)
		fmt.Println("Grade : ")
		if avg >= 80 {
			fmt.Println("Sangat baik")
		} else if avg >= 70 && avg < 80 {
			fmt.Println("Baik")
		} else {
			fmt.Println("Perlu belajar")
		}
		fmt.Println("")
	}
}

func sumAndAvgV2(siswa ...int) (int, int) {
	total := 0
	for _, key := range siswa {
		total += key
	}
	ratarata := total / len(siswa)
	return total, ratarata
}
