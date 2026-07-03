package database

var mysql string

func init() {
	mysql = "MySQL"
}
func GetDatabase() string {
	return mysql
}
