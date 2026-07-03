package main

import (
	"errors"
	"fmt"
)

type User struct {
	Username, Password string
}

var daftarUser []User

var ErrUsernameSudahAda = errors.New("username sudah dipakai")
var ErrUserTidakDitemukan = errors.New("username tidak ditemukan")
var ErrPasswordSalah = errors.New("password salah")

func Register(username string, password string) error {
	for _, u := range daftarUser {
		if u.Username == username {
			return ErrUsernameSudahAda
		}
	}
	daftarUser = append(daftarUser, User{Username: username, Password: password})
	return nil
}

func Login(username string, password string) error {
	for _, u := range daftarUser {
		if u.Username == username {
			if u.Password != password {
				return ErrPasswordSalah
			}
			return nil
		}
	}
	return ErrUserTidakDitemukan
}

func main() {
	err := Register("fuady", "fuady2006")
	if err != nil {
		fmt.Println("Gagal register.", err)
	} else {
		fmt.Println("Register sukses")
	}

	err = Login("fuady", "fuady2006")
	if err != nil {
		fmt.Println("Gagal login.", err)
	} else {
		fmt.Println("Login sukses")
	}
}
