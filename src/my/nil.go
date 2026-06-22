package main

import "fmt"

func CheckNil(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	result := CheckNil("")
	if result == nil {
		fmt.Println("Result is nil")
	} else {
		fmt.Println(result["name"])
	}
}
