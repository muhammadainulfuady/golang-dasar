package main

import "fmt"

func main() {
	books1 := []Buku{
		{
			"Laskar pelangi", "Indah", 100000,
		},
		{
			"Indah pelangi", "irwan", 50000,
		},
	}
	totalHarga := 0
	mahal := 0
	for _, key := range books1 {
		fmt.Println("Judul : ", key.Judul)
		fmt.Println("Harga : ", key.Harga)
		fmt.Println("Penulis : ", key.Penulis)
		fmt.Println("")
		totalHarga += key.Harga
		if key.Harga > mahal {
			mahal = key.Harga
		}
	}
	fmt.Println("Total harga semua buku : ", totalHarga)
	fmt.Println("Buku paling mahal : ", mahal)
}

type Buku struct {
	Judul   string
	Penulis string
	Harga   int
}

type Customer struct {
	Name, Addres string
	Age          int
}
