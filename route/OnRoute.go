package route

import (
	"github.com/gin-gonic/gin"
	"main.go/app/asms/controller"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	asms := router.Group("/asms")
	{
		asms.Use(func(context *gin.Context) {
		}, gin.Recovery())
		asms.Any("/", func(context *gin.Context) {
			context.String(0, asms.BasePath())
		})
		controller.IndexController(asms)
	}
}
