package main

import "fmt"

type Mahasiswaa struct {
	Nama  string
	Nilai int
}

func NaikkanNilai(m Mahasiswaa) {
	m.Nilai += 10
}

func main() {
	m := &Mahasiswaa{
		Nama:  "Budi",
		Nilai: 70,
	}

	fmt.Printf("Nlai %s sebelum naik = %d", m.Nama, m.Nilai)
	NaikkanNilai(m)
	fmt.Printf("\nNlai %s setelah naik = %d", m.Nama, m.Nilai)

}
