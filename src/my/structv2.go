package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Password string
	Phone    string
}

func main() {
	user := []User{
		{
			Name:     "Ainul Fuady",
			Email:    "ainul.fuady@example.com",
			Password: "password123",
			Phone:    "123-456-7890",
		},
		{
			Name:     "Rinastuti",
			Email:    "rinastuti@example.com",
			Password: "password456",
			Phone:    "098-765-4321",
		},
	}
	for _, u := range user {
		fmt.Println("\nUser :", u.Name)
		message, success := u.Login("ainul.fuady@example.com", "password123")
		if success {
			fmt.Println(message)
		} else {
			fmt.Println(message)
		}
	}
	newUser, success := register("John Doe", "john.doe@example.com", "password789", "909090090990")
	if success {
		user = append(user, newUser)
	}
	fmt.Println("")
	showUsers(user)
}

func showUsers(users []User) {
	for i, u := range users {
		fmt.Println("User", i+1)
		fmt.Println("Name:", u.Name)
		fmt.Println("Email:", u.Email)
		fmt.Println("Phone:", u.Phone)
		fmt.Println()
	}
}

func (u User) Login(email string, password string) (message string, success bool) {
	if u.Password == password && u.Email == email {
		return "Login successful. Welcome, " + u.Name + "!", true
	}
	return "Login failed. Email or password is incorrect. in user :" + u.Name, false
}

func register(name string, email string, password string, phone string) (User, bool) {
	fmt.Println("")
	if name == "" || email == "" || password == "" || phone == "" {
		fmt.Println("Name, email, password, and phone number are required for registration.")
		return User{}, false
	} else {
		fmt.Println("Registration successful for user:", name)
	}
	return User{
		Name:     name,
		Email:    email,
		Password: password,
		Phone:    phone,
	}, true
}
