package main

import (
	"backend/model"
	"backend/router"
	"fmt"
	"os/exec"
)

func runCommand() {
	cmd := exec.Command("swag", "init")
	fmt.Println("Cmd", cmd.Args)
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(out.String())
}

// @title           Golang-Admin-Boilerpalte Api
// @version         1.0
// @description     后端admin模版api
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
//// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	runCommand()
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
