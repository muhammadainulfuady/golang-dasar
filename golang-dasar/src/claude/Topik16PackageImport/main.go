package main

import (
	"belajar-golang-dasar/src/claude/Topik16PackageImport/helper"
	"belajar-golang-dasar/src/claude/Topik16PackageImport/math"
	"belajar-golang-dasar/src/claude/Topik16PackageImport/user"
	"fmt"
)

func init() {
	fmt.Println(helper.UcapanSelamatDatang())
}

func main() {
	hasil := math.Tambah(10, 5)

	fmt.Println(hasil)

	user1 := user.User{
		Nama: "ProgrammerHandal",
		Umur: 20,
	}

	user.TampilkanUser(user1)

}
