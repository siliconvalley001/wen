package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(router *gin.RouterGroup)  {
	router.POST("/login", UserLogin)
	router.POST("/register",UserRegister)
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"msg":"hellouser",
		})
	})

}
