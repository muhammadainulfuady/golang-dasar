package main

import "fmt"

func main() {
	tabungans := map[string][]int{
		"Budi":     {10000, 15000, 20000},
		"Andi":     {5000, 10000, 7000},
		"Siti":     {20000, 25000, 30000},
		"Ainul":    {12000, 15000, 13000},
		"Muhammad": {10000, 8000, 9000},
	}
	showData(tabungans)
	fmt.Println("")
	mainPrograms(tabungans)
}

// membuat fungsi untuk menampilkan data
func showData(tabungans map[string][]int) {
	for key, val := range tabungans {
		fmt.Println(key, ":", val)
	}
}

// membuat fungsi ibaratnya fungsi utama dari program ini
func mainPrograms(tabungans map[string][]int) {
	for key, val := range tabungans {
		total, avg := sumAllV2(val...)
		fmt.Println(key)
		fmt.Println("Total tabungan : ", total)
		fmt.Println("Rata rata tabungan : ", avg)
		if avg >= 20000 {
			fmt.Println("Status : Rajin Menabung")
		} else if avg >= 10000 {
			fmt.Println("Status : Cukup Rajin")
		} else {
			fmt.Println("Status : Perlu Menabung Lagi")
		}
		fmt.Println("")
	}
}

// membuat fungsi vardic karena ada banyak angka pada slice
func sumAllV2(number ...int) (int, int) {
	total := 0
	for _, value := range number {
		total += value
	}
	avg := total / len(number)
	return total, avg
}
