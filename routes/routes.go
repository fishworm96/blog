package routes

import (
	"blog/controller"
	"blog/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	// 判断模式
if mode == gin.ReleaseMode {
	gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signUp", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	return r
}