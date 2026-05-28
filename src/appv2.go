package main

import (
	"fmt"
	"math"
)

func main() {
	barangs := map[string]int{
		"Indomie": 10,
		"Susu":    3,
		"Roti":    7,
		"Gula":    0,
		"Kopi":    500,
		"Teh":     8,
		"Beras":   12,
	}
	daftarBarangs(barangs)
	fmt.Println("")
	statusPenjualan(barangs)
	fmt.Println("")
	sumAndAvg(barangs)
	fmt.Println("")
	barangTerlaris(barangs)
}

// fungsi ini dibuat untuk menampilkan daftar barang
func daftarBarangs(barang map[string]int) {
	for key, value := range barang {
		fmt.Println(key, ":", value)
	}
}

// status penjualan barang
func statusPenjualan(barang map[string]int) {
	for key, value := range barang {
		if value == 0 {
			fmt.Println(key, ": Tidak terjual")
		} else if value >= 1 && value <= 5 {
			fmt.Println(key, ": Penjualan rendah")
		} else {
			fmt.Println(key, ": Penjualan tinggi")
		}
	}
}

func sumAndAvg(barang map[string]int) {
	n := len(barang)
	total := 0
	for _, value := range barang {
		total += value
	}
	rataRata := total / n
	fmt.Println("Total penjualan : ", total)
	fmt.Println("Rata rata penjualan : ", rataRata)
}
func barangTerlaris(barang map[string]int) {
	topStok := math.MinInt
	lowStok := math.MaxInt
	topBarang := ""
	lowBarang := ""
	for key, value := range barang {
		if value > topStok {
			topStok = value
			topBarang = key
		}
		if value < lowStok {
			lowStok = value
			lowBarang = key
		}
	}
	fmt.Println("Stok paling sedikit :", lowBarang, "-", lowStok)
	fmt.Println("Stok paling banyak :", topBarang, "-", topStok)
}
