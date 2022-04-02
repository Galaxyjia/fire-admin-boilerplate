package model

type Article struct {
	ID        string `json:"id" example:"1"`
	FirstName string `json:"firstname" example:"1"`
	LastName  string `json:"lastname" example:"1"`
	Age       int    `json:"age" example:"1"`
}

func GetArticleList() {

}

func CreateArticle() {
	article := Article{ID: "aas", FirstName: "John", LastName: "asdf", Age: 32}
	DB.Create(&article)
}
