package router

import (
	"github.com/gin-gonic/gin"
	"mywork/controllers/account"
	"mywork/controllers/index"
	"mywork/middleware"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/resources", "./resources")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.GET("/", index.Index)

	router.GET("/jwt", index.Jwt)
	router.GET("/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	r1 := router.Group("/account").Use(middleware.SetHeaderJSON())
	{
		r1.GET("/", account.Func1)
		//r1.POST("/",account.Func2)
		//r1.POST("/a",account.Func3)
		//r1.POST("/b",account.Func4)
		r1.POST("/login", account.Login)
		r1.GET("/list", account.ListAccount)
		r1.POST("/reg", account.Register)
		r1.GET("/info/:id", account.Info)
		r1.DELETE("/rem/:id", account.Remove)
		r1.PUT("/update/:id", account.Update)
		r1.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
	}
	return router
}
