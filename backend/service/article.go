package service

import (
	"backend/model"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary articles list
// @Schemes
// @Description get user list
// @Tags Articles
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v1/api/articles [get]
func ArticleList(c *gin.Context) {
	model.GetArticleList()
	// model.CreateArticle()
	c.JSON(200, gin.H{
		"message": "list articles",
	})
}

// PingExample godoc
// @Summary user create
// @Schemes
// @Description create user list
// @Tags Articles
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v1/api/articles [post]
func ArticleCreate(c *gin.Context) {
	model.CreateArticle()
	c.JSON(200, gin.H{
		"message": "add user",
	})
}

// PingExample godoc
// @Summary user retrieve
// @Schemes
// @Description retrieve user
// @Tags Articles
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v1/api/articles/:id [get]
func ArticleRetrieve(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "retrieve user",
	})
}

// PingExample godoc
// @Summary user update
// @Schemes
// @Description update user
// @Tags Articles
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v1/api/articles/:id [put]
func ArticleUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update user",
	})
}

// PingExample godoc
// @Summary user delete
// @Schemes
// @Description delete user
// @Tags Articles
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v1/api/articles/:id [delete]
func ArticleDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete user",
	})
}
