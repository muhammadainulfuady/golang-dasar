package main

import (
	"belajar-golang-dasar/src/my/api"
	"belajar-golang-dasar/src/my/database"
	"belajar-golang-dasar/src/my/helper"
	_ "belajar-golang-dasar/src/my/internal"
	"fmt"
)

func main() {
	api := api.ApiCalling()
	result := helper.SayHello("Fuady maybe")
	database := database.GetDatabase()
	fmt.Println(result)
	fmt.Println(api)
	fmt.Println(database)
}
