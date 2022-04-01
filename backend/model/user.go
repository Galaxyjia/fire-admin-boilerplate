package model

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func GetUserList() {

}

func CreateUser() {
	user := User{ID: "aas", FirstName: "John", LastName: "asdf", Age: 32}
	DB.Create(&user)
}
