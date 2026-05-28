package main

import "fmt"

func message(messages string) string {
	return messages
}

func loop(n int) {
	for i := range n {
		fmt.Println(i)
	}
}

func main() {
	loop(10)
	defer fmt.Println(message("sudah berhenti saja anda"))
}
