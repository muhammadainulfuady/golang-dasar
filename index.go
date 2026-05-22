package main

func main() {
	biodata := []string{
		"address: 123 Main St",
		"age: 30",
		"email: johndoe@example.com",
	}
	for _, data := range biodata {
		println(data)
	}

	for i := 0; i < len(biodata); i++ {
		println(biodata[i])
	}
}
