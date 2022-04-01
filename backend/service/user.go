package service

import (
	"backend/model"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	model.GetUserList()
	// model.CreateUser()
	c.JSON(200, gin.H{
		"message": "list users",
	})
}

func UserCreate(c *gin.Context) {
	model.CreateUser()
	c.JSON(200, gin.H{
		"message": "add user",
	})
}

func UserUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update user",
	})
}

func UserDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete user",
	})
}
