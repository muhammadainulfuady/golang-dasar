package main

import "fmt"

func randomAssertion() any {
	return true
}

func main() {
	switch value := randomAssertion().(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Integer:", value)
	case bool:
		fmt.Println("Boolean:", value)
	default:
		fmt.Println("Unknown type")
	}
}