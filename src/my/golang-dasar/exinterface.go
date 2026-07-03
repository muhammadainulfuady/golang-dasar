package main

import "fmt"

func main() {
	data := map[string]any{
		"name":   "Muhammad Ainul Fuady",
		"umur":   20,
		"alamat": "Gresik",
	}

	for key, val := range data {
		fmt.Println(key, ":", val)
	}
}
