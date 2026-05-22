package main

import (
	"fmt"
	"strings"
)

func main() {
	result := tambah(5, 10)
	fmt.Println("Hasil:", result)
	fmt.Println(mbohlah())
	namadepan, namabelakang := namaLengkap("John", "Doe")
	fmt.Println("Nama Lengkap:", namadepan, namabelakang)
	namadepan2, _ := namaLengkap2()
	fmt.Println("Nama Lengkap 2:", namadepan2)
	firstname, middlename, lastname := namaLengkap3()
	fmt.Println("Nama Lengkap 3:", firstname, middlename, lastname)
	resultSummAll := sumAll(10, 10, 10, 10, 10)
	dataSlice := []int{
		10, 10, 10, 10, 10, 10,
	}
	fmt.Println(dataSlice)
	fmt.Println(resultSummAll)
	fmt.Println(sumAll(dataSlice...))

	getHello := getHelloWorld
	fmt.Println(getHello("adi"))

	message("Anjing", filterMessage)
}

func tambah(a int, b int) int {
	return a + b
}

func mbohlah() string {
	result := "Mbohlah"
	return result
}

// menggunakan parameter dan mengembalikan dua nilai
func namaLengkap(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

// tanpa parameter dan tanpa nilai kembali
func namaLengkap2() (string, string) {
	return "John", "Doe"
}

func namaLengkap3() (firstname, middlename, lastname string) {
	firstname = "Muhammad"
	return firstname, middlename, lastname
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getHelloWorld(name string) string {
	return "Hello World!! " + name
}

func message(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func filterMessage(message string) string {
	kataKotor := []string{
		"aNjinG",
		"kontol",
		"tai",
	}
	for _, filter := range kataKotor {
		if strings.ToLower(message) == strings.ToLower(filter) {
			return "****"
		}
	}
	return message
}
