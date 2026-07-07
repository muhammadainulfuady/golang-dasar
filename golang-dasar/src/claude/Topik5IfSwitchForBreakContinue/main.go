package main

import "fmt"

func main() {
	nilaiUjian := []int{80, 75, 90, 55, 40, 95}

	// TUGAS:
	// 1. Loop semua nilaiUjian pakai `for` + `range`
	// 2. Kalau nilai < 60, print "Nilai X: Tidak Lulus" (skip ke nilai berikutnya, jangan diproses lagi di bawah)
	// 3. Kalau nilai >= 90, print "Nilai X: Istimewa!" terus LANGSUNG STOP loop-nya (gak usah lanjut ngecek nilai setelahnya)
	// 4. Selain 2 kondisi di atas, print "Nilai X: Lulus"
	//
	// (X diganti sama nilai aslinya)

	for _, v := range nilaiUjian {
		if v < 60 {
			fmt.Printf("Nilai %d tidak lulus\n", v)
			continue
		} else if v >= 90 {
			fmt.Printf("Nilai %d istimewa!\n", v)
			break
		} else {
			fmt.Printf("Nilai %d Lulus\n", v)
		}
	}
}
