package myerror

import "errors"

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOl")
	} else {
		return nilai / pembagi, nil
	}
}
