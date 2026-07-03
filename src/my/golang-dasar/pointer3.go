package main

import "fmt"

type Siswa struct {
	Name   string
	Nisn   string
	Alamat string
}

func updateSiswa(s *Siswa, newName string, newNisn string, newAlamat string) {
	s.Name = newName
	s.Nisn = newNisn
	s.Alamat = newAlamat
}

func tambah(a, b int, hasil *int) {
	*hasil = a + b
}

func main() {
	s := &Siswa{
		Name:   "Muhammad Ainul Fuady",
		Nisn:   "00678765",
		Alamat: "Pereng kulon, melirang, bungah, gresik, jawa timur",
	}
	fmt.Println(s)
	updateSiswa(s, "Irfan", "123321", "nothig")
	fmt.Println(s)
	var hasil int
	tambah(1, 1, &hasil)
	fmt.Println(hasil)
}
