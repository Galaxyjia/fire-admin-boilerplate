package main

import (
	"backend/model"
	"backend/router"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

// var DB *gorm.DB

func main() {
	// 全局DB,注意不能用 DB,err := (这里为局部变量)
	// DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// Migrate the schema
	//自动迁移数据库
	// DB.AutoMigrate(&User{})

	// r := gin.Default()
	// r.GET("/ping/:id", pingGet)
	// r.POST("/ping", pingCreate)

	// 初始化DB 数据库
	model.Init()

	// 初始化路由
	r := router.Router()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
