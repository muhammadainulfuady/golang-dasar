package main

import (
	"belajar-golang-dasar/src/my/api"
	"belajar-golang-dasar/src/my/database"
	"belajar-golang-dasar/src/my/helper"
	_ "belajar-golang-dasar/src/my/internal"
	"belajar-golang-dasar/src/my/myerror"
	"fmt"
)

func main() {
	api := api.ApiCalling()
	result := helper.SayHello("Fuady maybe")
	database := database.GetDatabase()
	fmt.Println(result)
	fmt.Println(api)
	fmt.Println(database)

	hasil, err := myerror.Pembagian(100, 0)
	if err == nil {
		fmt.Printf("Hasil = %d", hasil)
	} else {
		fmt.Printf("Error = %s", err)
	}
}
