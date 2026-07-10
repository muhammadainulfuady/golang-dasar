package main

import (
	"belajar-golang-dasar/src/claude/Topik16PackageImport/math"
	"fmt"
)

func main() {
	tambahFromMath := math.Tambah(10, 3)
	kaliFromMath := math.Kali(10, 3)
	fmt.Println(tambahFromMath)
	fmt.Println(kaliFromMath)
}
