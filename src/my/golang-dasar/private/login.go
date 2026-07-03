package main


type Student struct {
	FullName, Username, Password string
}

func login(fullname, username, password string, s *Student) {
	s.FullName = fullname
	s.Username = username
	s.Password = password
}

func main() {
	login("a", "s", "u", &Student{
		FullName: "Indah",
		Username: "Indah123",
		Password: "indah",
	})
}
