package main

import "fmt"

type Kendaraan interface {
	Jalan()
}

type Mobil struct {
	Merk string
}

type Motor struct {
	Merk string
}

type Sepeda struct {
	Merk string
}

func (m Mobil) Jalan() {
	fmt.Println(m.Merk, "Sedang berjalan")
}

func (m Motor) Jalan() {
	fmt.Println(m.Merk, "Sedang berjalan")
}

func (s Sepeda) Jalan() {
	fmt.Println(s.Merk, "Sedang berjalan")
}

func tampilkanJalan(k Kendaraan) {
	k.Jalan()
}

func main(){
	mb1 := Mobil{
		Merk: "Avanza",
	}

	mt1 := Motor{
		Merk: "Beat",
	}

	sp1 := Sepeda{
		Merk: "Polygon",
	}
	tampilkanJalan(mb1)
	tampilkanJalan(mt1)
	tampilkanJalan(sp1)
}
