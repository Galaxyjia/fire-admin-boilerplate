package service

import (
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary users list
// @Schemes
// @Description get user list
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /admin/users [get]
func UserList(c *gin.Context) {
	model.GetUserList()
	c.JSON(200, gin.H{
		"code":    200,
		"message": "list users",
	})
}

// PingExample godoc
// @Summary user create
// @Schemes
// @Description create user list
// @Tags Users
// @Accept json
// @Produce json
// @Param       data	body	string  true  "Request Data" SchemaExample({\r\n"id":"kkkkasdf",\r\n"firstname":"gaalaxy",\r\n"lastname":"fasdfa",\r\n"age":12\r\n})
// @Success 200 {object} model.User
// @Router /admin/users [post]
func UserCreate(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.CreateUser(user.ID, user.FirstName, user.LastName, user.Age)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "add user",
	})
}

// PingExample 	godoc
// @Summary 	user retrieve
// @Schemes
// @Description retrieve user
// @Tags 		Users
// @Accept 		json
// @Produce 	json
// @Param       id	path	string  true  "Account ID"
// @Success 	200 {string} Helloworld
// @Router 		/admin/users/{id} [get]
func UserRetrieve(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "retrieve user",
		"data": gin.H{
			"id": id,
		},
	})
}

// PingExample godoc
// @Summary user update
// @Schemes
// @Description update user
// @Tags Users
// @Accept json
// @Produce json
// @Param       id	path	string  true  "Account ID"
// @Success 200 {string} Helloworld
// @Router /admin/users/{id} [put]
func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "update user",
		"data": gin.H{
			"id": id,
		},
	})
}

// PingExample godoc
// @Summary user delete
// @Schemes
// @Description delete user
// @Tags Users
// @Accept json
// @Produce json
// @Param       id	path	string  true  "Account ID"
// @Success 200 {string} Helloworld
// @Router /admin/users/{id} [delete]
func UserDelete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "delete user",
		"data": gin.H{
			"id": id,
		},
	})
}
