package main

import "fmt"

type Kucing struct {
	Name string
}

type Anjing struct {
	Name string
}

type Hewan interface {
	Bersuara()
}

type Payment interface {
	Pay()
	Print()
}

func (k Kucing) Bersuara() {
	fmt.Println("Meong meong")
}
func (a Anjing) Bersuara() {
	fmt.Println("Guk guk")
}
func suruhBersuara(h Hewan) {
	h.Bersuara()
}

func main() {
	kucing := Kucing{
		Name: "Kucing lucu banget",
	}
	anjing := Anjing{
		Name: "Kucing lucu banget",
	}
	suruhBersuara(kucing)
	suruhBersuara(anjing)
	kucing.Bersuara()
	anjing.Bersuara()
}
