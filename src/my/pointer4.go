package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr." + man.Name
	fmt.Println(man.Name)
}

func main() {
	budi := &Man{
		Name: "Budi",
	}
	fmt.Println(budi.Name)
	// fmt.Printf("%s sebelum menikah %s", budi.Name, budi.Name)
	budi.Married()
	// fmt.Printf("\n%s setelah menikah %s", budi.Name, budi.Name)
}
