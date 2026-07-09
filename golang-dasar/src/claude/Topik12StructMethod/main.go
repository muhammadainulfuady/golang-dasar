package main

import "fmt"

type RekeningBank struct {
	NamaPemilik string
	Saldo       int
}

func (r *RekeningBank) Setor(jumlah int) {
	r.Saldo += jumlah
}

func (r *RekeningBank) Tarik(jumlah int) {
	if r.Saldo < jumlah {
		fmt.Println("Saldo tidak cukup ya")
		return
	}
	r.Saldo -= jumlah
}

func (r *RekeningBank) CekSaldo() int {
	return r.Saldo
}


func main(){
	rekening1 := RekeningBank{
		NamaPemilik: "Saya sendiri",
		Saldo: 20000,
	}
	rekening1.Setor(10000)
	rekening1.Tarik(5000)
	fmt.Println(rekening1.CekSaldo())
}