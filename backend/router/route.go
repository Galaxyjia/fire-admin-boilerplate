package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
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

	// r.GET("json", func(c *gin.Context) {
	// 	// 方法1：使用map
	// 	// data := map[string]interface{}{
	// 	// 	"data": map[string]interface{}{
	// 	// 		"name":    "galaxy",
	// 	// 		"age":     18,
	// 	// 	},
	// 	// 	"msg":  "返回成功",
	// 	// 	"code": 200,
	// 	// }

	// 	// 方法2: gin.H
	// 	// data := gin.H{
	// 	// 	"data": gin.H{
	// 	// 		"name": "galaxy",
	// 	// 		"age":  18,
	// 	// 	},
	// 	// 	"msg":  "返回成功",
	// 	// 	"code": 200,
	// 	// }

	// 	// 方法3: 结构体

	// 	data := User{
	// 		FirstName: "galaxy",
	// 		LastName:  "guo",
	// 		Age:       18,
	// 	}

	// 	c.JSON(http.StatusOK, data)
	// })

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
		apiRouters.GET("/search", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "search list",
			})
		})

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

	return r
}

// func pingGet(c *gin.Context) {
// 	id := c.Param("id")
// 	firstname := c.DefaultQuery("firstname", "Guest")
// 	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
// 	user := User{ID: id, FirstName: firstname, Age: 18, LastName: lastname}
// 	DB.Create(&user)

// 	c.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "pong",
// 		"data": gin.H{
// 			"id":        id,
// 			"firstname": firstname,
// 			"lastname":  lastname,
// 		},
// 	})
// }

// func pingCreate(c *gin.Context) {
// 	// id := c.Param("id")
// 	// firstname := c.DefaultQuery("firstname", "Guest")
// 	// lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
// 	// user := User{ID: id, FirstName: firstname, Age: 18, LastName: lastname}
// 	// DB.Create(&user)
// 	var user User

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	DB.Create(User{ID: user.ID, FirstName: user.FirstName, Age: user.Age, LastName: user.LastName})

// 	c.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "pong",
// 		"data": gin.H{
// 			"id":        user.ID,
// 			"firstname": user.FirstName,
// 			"lastname":  user.LastName,
// 		},
// 	})
// }
