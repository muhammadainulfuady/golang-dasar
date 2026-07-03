// package main

// import "fmt"

// type Kendaraan interface {
// 	Jalan() string
// 	BahanBakar() string
// }

// type Mobil struct {
// 	Merk string
// }

// type Motor struct {
// 	Merk string
// }

// type Sepeda struct {
// 	Merk string
// }

// func (m Mobil) Jalan() string {
// 	return m.Merk + " sedang berjalan"
// }

// func (m Mobil) BahanBakar() string {
// 	return "Bensin"
// }

// func (m Motor) Jalan() string {
// 	return m.Merk + " sedang berjalan"
// }

// func (m Motor) BahanBakar() string {
// 	return "Pertalite"
// }

// func (s Sepeda) Jalan() string {
// 	return s.Merk + " sedang berjalan"
// }

// func (s Sepeda) BahanBakar() string {
// 	return "Tenaga manusia"
// }

// func tampilkanJalan(k Kendaraan) {
// 	fmt.Println(k.Jalan())
// 	fmt.Println(k.BahanBakar())
// }

// func main(){
// 	mb1 := Mobil{
// 		Merk: "Avanza",
// 	}

// 	mt1 := Motor{
// 		Merk: "Beat",
// 	}

// 	sp1 := Sepeda{
// 		Merk: "Polygon",
// 	}
// 	tampilkanJalan(mb1)
// 	tampilkanJalan(mt1)
// 	tampilkanJalan(sp1)
// }

package main

import "fmt"

type Kendaraan interface {
	Jalan() string
	Info() string
	BahanBakarV1() string
}

type Mobil struct {
	Merk       string
	BahanBakar string
}

type Motor struct {
	Merk       string
	BahanBakar string
}

type Sepeda struct {
	Merk       string
	BahanBakar string
}

func (m Mobil) Jalan() string {
	return m.Merk + " sedang berjalan"
}

func (m Mobil) BahanBakarV1() string {
	return "Menggunakan bahan bakar: " + m.BahanBakar
}
func (m Mobil) Info() string {
	return "Mobil"
}

func (m Motor) Jalan() string {
	return m.Merk + " sedang berjalan"
}

func (m Motor) BahanBakarV1() string {
	return "Menggunakan bahan bakar: " + m.BahanBakar
}

func (m Motor) Info() string {
	return "Motor"
}

func (s Sepeda) Jalan() string {
	return s.Merk + " sedang berjalan"
}

func (s Sepeda) BahanBakarV1() string {
	return "Menggunakan bahan bakar: " + s.BahanBakar
}

func (s Sepeda) Info() string {
	return "Sepeda"
}

func tampilkanJalan(k Kendaraan) {
	fmt.Println("\nJenis Kendaraan:", k.Info())
	fmt.Println(k.Jalan())
	fmt.Println(k.BahanBakarV1())
}

func main() {
	mb1 := Mobil{
		Merk:       "Avanza",
		BahanBakar: "Bensin",
	}

	mt1 := Motor{
		Merk:       "Beat",
		BahanBakar: "Pertalite",
	}

	sp1 := Sepeda{
		Merk:       "Polygon",
		BahanBakar: "Tenaga manusia",
	}
	tampilkanJalan(mb1)
	tampilkanJalan(mt1)
	tampilkanJalan(sp1)
}
