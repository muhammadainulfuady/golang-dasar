package main

import "fmt"

func tampilkan(data any) {
	// TUGAS:
	// Gunakan TYPE SWITCH.
	switch v := data.(type) {

	case string:
		// Kalau string:
		// "Ini String: <isi>"
		fmt.Println("Ini String : ", v)
	case int:
		// Kalau int:
		// "Ini Integer: <isi>"
		fmt.Println("Ini Integer : ", v)
	case bool:
		// Kalau bool:
		// "Ini Boolean: <isi>"
		fmt.Println("Ini Boolean : ", v)
	default:
		// Selain itu:
		// "Tipe tidak dikenali"
		fmt.Println("Tipe tidak dikenali")
	}
}

func main() {
	tampilkan("Halo")
	tampilkan(100)
	tampilkan(true)
	tampilkan(3.14)
}
