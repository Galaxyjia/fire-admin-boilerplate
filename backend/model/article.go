package model

type Article struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func GetArticleList() {

}

func CreateArticle() {
	article := Article{ID: "aas", FirstName: "John", LastName: "asdf", Age: 32}
	DB.Create(&article)
}
