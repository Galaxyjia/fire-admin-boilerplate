package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "ok",
		"msg":  "pong",
	})

}

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
