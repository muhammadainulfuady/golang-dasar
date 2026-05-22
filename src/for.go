package main

func main() {

	// seperti while loop pada bahasa pemrograman lain
	number := 1
	for number <= 5 {
		println("jancok ke-", number)
		number++
	}

	// seperti for loop pada bahasa pemrograman lain
	for i := 1; i <= 5; i++ {
		println("jancok ke-", i)
	}

	// for range
	seseorang := []string{
		"indra",
		"budi",
		"andi",
		"siti",
		"rudi",
		"sari",
		"adit",
		"adi",
	}

	for _, nama := range seseorang {
		println("nama:", nama)
	}	

	// atau for len biasa
	for i := 0; i < len(seseorang); i++ {
		println("nama:", seseorang[i])
	}


	// menggunakan index value
	for index, nama := range seseorang {
		if index%2 == 0 {
			println("index genap:", index, "nama:", nama)
		} else {
			println("index ganjil:", index, "nama:", nama)
		}
	}
}

// jancok-1
// jancok-2
// jancok-3
// jancok-4
// jancok-5
