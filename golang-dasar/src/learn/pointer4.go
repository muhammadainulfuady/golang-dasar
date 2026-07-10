package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	budi := Man{
		Name: "Budi",
	}
	fmt.Printf("%s sebelum menikah %s", budi.Name, budi.Name)
	budi.Married()
	fmt.Printf("\n%s setelah menikah %s", budi.Name, budi.Name)
}
