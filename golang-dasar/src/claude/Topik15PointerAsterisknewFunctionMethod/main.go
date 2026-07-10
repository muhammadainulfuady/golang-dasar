package main

import (
	"errors"
	"fmt"
)

var (
	ErrSaldoTidakCukup = errors.New("Saldo tida cukup")
)

type AtmMini struct {
	NamaAtm     string
	NamaNasabah string
	Saldo       int
}

type RepositoryAtmMini interface {
	Tarik(jumlah int) (int, error)
	Tambah(jumlah int)
}

func (a *AtmMini) Tarik(jumlah int) (int, error) {
	if a.Saldo < jumlah {
		return a.Saldo, ErrSaldoTidakCukup
	}
	a.Saldo -= jumlah
	return a.Saldo, nil
}

func (a *AtmMini) Tambah(jumlah int) {
	a.Saldo += jumlah
}

func MenarikUang(r RepositoryAtmMini, jumlah int) (int, error) {
	return r.Tarik(jumlah)
}

func MenambahkUang(r RepositoryAtmMini, jumlah int) {
	r.Tambah(jumlah)
}

func main() {
	BRI := &AtmMini{
		NamaAtm:     "BRI",
		NamaNasabah: "Irwan",
		Saldo:       20000,
	}

	fmt.Printf("SALDO AWAL : %d\n", BRI.Saldo)
	hasilNarik, err := MenarikUang(BRI, 13000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hasilNarik)
	}

}
