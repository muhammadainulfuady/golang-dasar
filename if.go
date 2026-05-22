package main

func main() {
	biodata := []string{
		"address: 123 Main St",
		"age: 30",
		"email: johndoe@example.com",
	}

	/*
		if expression dengan statement pendek karena terdapat variabel yang hanya digunakan untuk if expression tersebut.
	*/
	if lenght := len(biodata); lenght >= 3 {
		println("Biodata has more than 3 items")
	}

	// if expression biasa
	age := 17
	if age >= 18 {
		println("anda bisa membuka ktp")
	} else {
		println("anda belum bisa membuka ktp")
	}
}
