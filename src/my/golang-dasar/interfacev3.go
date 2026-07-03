package main

import "fmt"

type Notifikasi interface {
	KirimNotifikasi(pesan string)
}

type Email struct {
	AlamatEmail string
}

func (e Email) KirimNotifikasi(pesan string) {
	fmt.Print(e.AlamatEmail, "Send with message:\n", pesan)
}

func KirimNotifikasi(n Notifikasi, pesan string) {
	n.KirimNotifikasi(pesan)
}

func main() {
	email := Email{
		AlamatEmail: "ainulfuadi@gmail.com",
	}
	KirimNotifikasi(email, "Hello my name is Ainul Fuadi, i am learning Golang programming language")
}
