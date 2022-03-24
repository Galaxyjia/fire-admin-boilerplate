# fire-admin-boilerplate

cd backend

## 初始化go项目
go mod init fire-admin-boilerplate

## 创建main.go
```
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

## 执行 go mod tidy 加载模块

## 执行 go run main.go 
访问 localhost:8080/ping


## json返回
```
r.GET("json", func(c *gin.Context) {
		// 方法1：使用map
		data := map[string]interface{}{
			"data": map[string]interface{}{
				"name":    "galaxy",
				"message": "hello world",
				"age":     18,
			},
			"msg":  "返回成功",
			"code": 200,
		}

		c.JSON(http.StatusOK, data)
	})
```
