package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "New York",
		Province: "NY",
		Country:  "USA",
	}

	// pass by value, so address2 is a copy of address1
	address2 := address1
	address2.City = "Los Angeles"
	fmt.Println(address1)
	fmt.Println(address2)

	// pass by reference, so address3 points to the same memory location as address1
	address3 := &address1
	address3.City = "Chicago"
	fmt.Println(address1)
	fmt.Println(address3)

}
