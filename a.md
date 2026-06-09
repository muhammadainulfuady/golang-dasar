
Materi Video :
00:00:00 - Pendahuluan
00:02:41 - Pengenalan Golang
00:07:53 - Menginstall Golang
00:11:46 - Membuat Project
00:15:35 - Program Hello World
00:26:11 - Multiple Main Function
00:30:29 - Tipe Data Number
00:38:14 - Tipe Data Boolean
00:40:12 - Tipe Data String
00:45:23 - Variable
00:54:37 - Constant
00:58:27 - Konversi Tipe Data
01:05:49 - Type Declarations
01:09:08 - Operasi Matematika
01:14:51 - Operasi Perbandingan
01:21:34 - Operasi Boolean
01:27:46 - Tipe Data Array
01:40:02 - Tipe Data Slice
02:07:16 - Tipe Data Map
02:14:26 - If
02:25:00 - Switch
02:32:13 - For
02:41:52 - Break dan Continue
02:46:11 - Function
02:49:12 - Function Parameter
02:53:35 - Function Return Value
02:58:02 - Returning Multiple Values
03:02:50 - Named Return Values
03:06:33 - Variadic Function
03:13:51 - Function as Value
03:17:30 - Function as Parameter
03:25:29 - Anonymous Function
03:30:50 - Recursive Function
03:37:14 - Closure
03:41:20 - Defer Panic dan Recover
03:51:50 - Komentar


saya sudah belajar mengenai bahasa pemograman golang seperti yang tercantum pada materi di atas, nah setelah saya belajar, maka alangkah baiknay berikan saya 1 soal tetapi kompleks seperti membuat sebuah aplikasi, tetapi hanya pada materi yang ada pada diatas jadi yang belum saya pelajari diatas itu tidak perlu anda berikan seperti itu, yah tidak perlu menggunakan semua materi yang ada pada diatas karena akan saya gunakan logika saya sepenuhnya, nanti anda berikan soalnya saja buatkan aplikasi sebagai berikut... ngerti yah?


func main() {
	user := []User{
		{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
			Phone:    "123-456-7890",
		},
		{
			Name:     "Jane Smith",
			Email:    "jane.smith@example.com",
			Password: "password456",
			Phone:    "098-765-4321",
		},
	}
	for _, u := range user {
		if u.Login("john.doe@example.com", "password123") {
			fmt.Println("Login successful for user:", u.Name)
		} else {
			fmt.Println("Login failed for user:", u.Name)
		}
	}
}
func (u User) Login(email string, password string) bool {
	if u.Email == email && u.Password == password {
		return true
	}
	return false
}

saya  mau nanya sih lebih fleksibel structnya sendiri2 apa dari slice gitu? soalnya kalau dijadikan slice bakalan isi Login kan 1 kali yah , jadi kalau menurut saya mending di pisah tidak dijadikan seperti slice gitu, cuman yah saya harus buat user1, user2, user3 dst..