package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping)

	r.GET("json", func(c *gin.Context) {
		// 方法1：使用map
		// data := map[string]interface{}{
		// 	"data": map[string]interface{}{
		// 		"name":    "galaxy",
		// 		"message": "hello world",
		// 		"age":     18,
		// 	},
		// 	"msg":  "返回成功",
		// 	"code": 200,
		// }

		// 方法2: gin.H
		data := gin.H{
			"data": gin.H{
				"name":    "galaxy",
				"message": "hello world",
				"age":     18,
			},
			"msg":  "返回成功",
			"code": 200,
		}

		// 方法3: 结构体

		c.JSON(http.StatusOK, data)
	})

	r.GET("/routers", func(c *gin.Context) {
		routers := r.Routes()
		l := []string{}
		for _, v := range routers {
			fmt.Println(v.Method)
			fmt.Println(v.Path)
			l = append(l, v.Method+" : "+v.Path)
		}
		fmt.Println(len(l))
		fmt.Println(l)
		c.JSON(200, gin.H{
			"path": l,
		})
	})
	// 路由组
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "users",
			})
		})
		adminRouters.POST("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "add users",
			})
		})
		adminRouters.PUT("/users/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "put users 1",
			})
		})
		adminRouters.DELETE("/users/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "delete users 1",
			})
		})
	}

	apiRouters := r.Group("/api/v1")
	{
		apiRouters.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "users list",
			})
		})
		apiRouters.POST("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "add users",
			})
		})

		apiRouters.GET("/users/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "get users 1",
			})
		})

		apiRouters.PUT("/users/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "put users 1",
			})
		})
		apiRouters.DELETE("/users/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "delete users 1",
			})
		})

		apiRouters.GET("/articles", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "articles list",
			})
		})
		apiRouters.POST("/articles", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "add articles",
			})
		})

		apiRouters.GET("/articles/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "get articles 1",
			})
		})

		apiRouters.PUT("/articles/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "put articles 1",
			})
		})
		apiRouters.DELETE("/articles/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "delete articles 1",
			})
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
