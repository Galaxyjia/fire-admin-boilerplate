package router

import (
	"backend/service"
	"fmt"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// docs.SwaggerInfo.BasePath = "/api/v1"

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 获取全部路由
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
		adminRouters.GET("/users", service.UserList)
		adminRouters.POST("/users", service.UserCreate)
		adminRouters.GET("/users/:id", service.UserRetrieve)
		adminRouters.PUT("/users/:id", service.UserUpdate)
		adminRouters.DELETE("/users/:id", service.UserDelete)
	}

	apiRouters := r.Group("/api/v1")
	{
		apiRouters.GET("/search", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "search list",
			})
		})

		apiRouters.GET("/articles", service.ArticleList)
		apiRouters.POST("/articles", service.ArticleCreate)
		apiRouters.GET("/articles/1", service.ArticleRetrieve)
		apiRouters.PUT("/articles/1", service.ArticleUpdate)
		apiRouters.DELETE("/articles/1", service.ArticleDelete)
	}

	return r
}
