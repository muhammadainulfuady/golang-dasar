package main

import "fmt"

type Nasabah struct {
	FullName string
	Atm      string
	Saldo    int
}

type ErrPenarikanUang struct {
	FullName string
	Atm      string
	Saldo    int
}

func (e ErrPenarikanUang) Error() string {
	return "Gagal narik saldo"
}

func (n *Nasabah) MenarikUang(jumlah int) (int, error) {
	if n.Saldo < jumlah {
		return 0, &ErrPenarikanUang{
			FullName: n.FullName,
			Saldo:    n.Saldo,
			Atm:      n.Atm,
		}
	}
	n.Saldo -= jumlah
	return n.Saldo, nil
}

func main() {
	nasabah1 := &Nasabah{
		FullName: "Nama Lengkap Saya",
		Atm:      "BRI",
		Saldo:    20000,
	}
	result, err := nasabah1.MenarikUang(2000000)
	if err != nil {
		if narikErr, ok := err.(*ErrPenarikanUang); ok {
			fmt.Println("Silahkan hubungi admin via chat yang biodata dibawah ini :")
			fmt.Println(narikErr.Atm)
			fmt.Println(narikErr.FullName)
			fmt.Println(narikErr.Saldo)
			return
		}
	} else {
		fmt.Println("Berhasil nari sisa salfo", result)
	}
}
