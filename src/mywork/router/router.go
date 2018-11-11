package router

import (
	"github.com/gin-gonic/gin"
	"mywork/controllers/account"
	"mywork/controllers/index"
	"mywork/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/favicon.ico")
	router.GET("/", index.Index)

	router.GET("/jwt", index.Jwt)

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
	}
	return router
}
