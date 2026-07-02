package main

import (
	"errors"
	"fmt"
)

func belaJaMakan() (string, error) {
	return "", errors.New("warungnya tutup")
}

func main() {
	result, err := belaJaMakan()
	if err != nil {
		fmt.Println("gagal", err)
	} else {
		fmt.Println("sukses", result)
	}
}
