package main

import "fmt"

type Rekening struct {
	Pemilik string
	Saldo   int
}

func (r Rekening) LihatSaldo() int {
	return r.Saldo
}

func (r *Rekening) Setor(jumlah int) {
	r.Saldo += jumlah
}

func (r *Rekening) Tarik(jumlah int) bool {
	if r.Saldo < jumlah {
		return false
	}

	r.Saldo -= jumlah
	return true
}

func Transfer(dari, ke *Rekening, jumlah int) bool {
	if !dari.Tarik(jumlah) {
		return false
	}
	ke.Setor(jumlah)
	return true
}

func main() {
	r1 := &Rekening{
		Pemilik: "Muhammad Ainul Fuady",
		Saldo:   5000000,
	}
	fmt.Printf("Saldo awal adalah %d", r1.Saldo)
	r1.Setor(1000000)
	fmt.Printf("Saldo setelah di tambah %d", r1.Saldo)
	tarik := r1.Tarik(2500000)
	if tarik == true {
		fmt.Printf("Saldo setelah di tarik %d", r1.Saldo)
	} else {
		fmt.Println("Maaf saldo yang anda tarik itu kurang")
	}
	Transfer(r1, r1, 1000000)
	fmt.Printf("Saldo setelah di transfer %d", r1.Saldo)
	fmt.Scanln("hell")
}
