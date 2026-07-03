package main

import "fmt"

type Hewan interface {
	Bersuara() string
}

type Kucing struct {
	Nama string
}

type Anjing struct {
	Nama string
}

func (k Kucing) Bersuara() string {
	return k.Nama + ": Meow..."
}

func (a Anjing) Bersuara() string {
	return a.Nama + ": Guk Guk..."
}

func tampilkanSuara(h Hewan) {
	fmt.Println(h.Bersuara())
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
