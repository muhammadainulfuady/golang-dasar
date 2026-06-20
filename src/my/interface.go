package main

import "fmt"

type Hewan interface {
	Bersuara()
}

type Kucing struct {
	Nama string
}

type Anjing struct {
	Nama string
}

func (k Kucing) Bersuara() {
	fmt.Println(k.Nama, ": Meong...")
}

func (a Anjing) Bersuara() {
	fmt.Println(a.Nama, ": Guk Guk...")
}

func tampilkanSuara(h Hewan) {
	h.Bersuara()
}

func main() {
	k1 := Kucing{
		Nama: "Kitty",
	}
	a1 := Anjing{
		Nama: "Doggy",
	}

	tampilkanSuara(k1)
	tampilkanSuara(a1)
}
