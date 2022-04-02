package model

type User struct {
	ID        string `json:"id" example:"1"`
	FirstName string `json:"firstname" example:"1"`
	LastName  string `json:"lastname" example:"1"`
	Age       int    `json:"age" example:"1"`
}

func GetUserList() {

}

func CreateUser(id string, firstname string, lastname string, age int) {
	user := User{ID: id, FirstName: firstname, LastName: lastname, Age: age}
	DB.Create(&user)
}
