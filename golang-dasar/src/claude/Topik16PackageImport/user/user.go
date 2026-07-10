package user

import "fmt"

type User struct {
	Nama string
	Umur int
}

func TampilkanUser(u User) {
	fmt.Printf("Nama : %s\n", u.Nama)
	fmt.Printf("Umur : %d\n", u.Umur)
}
