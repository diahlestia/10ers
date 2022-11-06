package routes

import (
	"10xers/controllers"

	"github.com/gin-gonic/gin"
)

func TokenRoute(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	router.GET("public/:wallet_address/contents", controllers.Get())
	router.POST("public/fetch_tokens", controllers.Create())
	router.DELETE("public/token/delete", controllers.Delete())
}
