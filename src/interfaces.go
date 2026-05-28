// sebenarnya field yag terdapat pada interface HashName itu untuk apa?
// package main

// import "fmt"

// type HashName interface {
// 	Getname() string
// }

// func SayHello(value HashName) {
// 	fmt.Println("Hello", value.Getname())
// }

// type Person struct {
// 	Name string
// }

// func (person Person) Getname() string {
// 	return person.Name
// }

// func main() {
// 	ps := Person{
// 		Name: "adi",
// 	}
// 	SayHello(ps)
// }

// package main

// import "fmt"

// type HargaBooks interface {
// 	Getprice() int
// }

// func SayHello(value HargaBooks) {
// 	fmt.Println("Hello", value.Getprice())
// }

// type Books struct {
// 	Name    string
// 	Harga   int
// 	Halaman int
// }

// func (p Books) Getprice() int {
// 	return p.Harga
// }

// func main() {
// 	b := Books{
// 		Name: "adi",
// 		Harga: 100000,
// 		Halaman: 120,
// 	}
// 	SayHello(b)
// }

// package main

// import "fmt"

// type Kendaraan interface {
// 	Jalan()
// }

// type Mobil struct {
// 	Nama      string
// 	Kecepatan int
// }

// type Motor struct {
// 	Nama      string
// 	Kecepatan int
// }

// func (m Mobil) Jalan() {
// 	fmt.Println("Mobil", m.Nama, "berjalan dengan kecepatan", m.Kecepatan)
// }

// func (m Motor) Jalan() {
// 	fmt.Println("Motor", m.Nama, "berjalan dengan kecepatan", m.Kecepatan)
// }

// func jalan(k Kendaraan) {
// 	k.Jalan()
// }

// func main() {
// 	mb1 := Mobil{
// 		Nama:      "Avanza",
// 		Kecepatan: 120,
// 	}
// 	mt1 := Motor{
// 		Nama:      "Beat",
// 		Kecepatan: 80,
// 	}
// 	jalan(mb1)
// 	jalan(mt1)
// }
