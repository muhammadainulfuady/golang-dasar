package main

import "fmt"


func main() {
	dataMahasiswa := map[string]any{
		"nama": "John Doe",
		"usia": 20,
		"jurusan": "Informatika",
		"nilai": 3.5,
	}

	for index, value := range dataMahasiswa {
		fmt.Println(index, value)
	}
}