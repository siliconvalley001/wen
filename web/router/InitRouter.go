package router

import (
	"github.com/gin-gonic/gin"
	gin_Swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_"github.com/siliconvalley001/wen/web/docs"
	"github.com/siliconvalley001/wen/web/middlemare"

	"github.com/siliconvalley001/wen/web/controller/user"

)

var url=gin_Swagger.URL("http://127.0.0.1:8089/swagger/doc.json")
////Vue.prototype.$http = axios
func InitRouter(router *gin.Engine)*gin.Engine{
	router.GET("/swagger/*any",gin_Swagger.WrapHandler(swaggerFiles.Handler,url))
	router.Use(middlemare.CrosMiddleWare)
	userGroup:=router.Group("/user")
	user.Router(userGroup)
	return router
}
