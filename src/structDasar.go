package main

import "fmt"

func main() {
	joko := Customer{
		Name:   "Joko",
		Addres: "Indonesia",
		Age:    20,
	}
	imam := Customer{
		Name:   "Imam",
		Addres: "Jepang",
		Age:    27,
	}
	ikul := Customer{
		Name:   "Ikul",
		Addres: "Selandia baru",
		Age:    17,
	}
	joko.sayHello("Agus")
	ikul.sayHello("Agus")
	imam.sayHello("Agus")
}

func (message Customer) sayHello(name string) {
	fmt.Println("Halo", name, "my name is", message.Name)
}

type Customer struct {
	Name, Addres string
	Age          int
}
