package routes

import (
	"blog/controller"
	"blog/logger"
	"blog/middlewares"

	_ "blog/docs" // 千万不要忘了导入把你上一步生成的docs

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	// 判断模式
if mode == gin.ReleaseMode {
	gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")

	v1.POST("/signUp", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/post", controller.GetPostListHandler)
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.GET("/tag", controller.GetTagListHandler)
	
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.CreatePostHandler)
		v1.PUT("/post/edit/:id", controller.UpdatePostHandler)
		v1.DELETE("/post/delete/:id", controller.DeletePostHandler)
		v1.POST("/vote", controller.PostVoteController)
	}
	{
		v1.POST("/tag", controller.CreateTagHandler)
		v1.PUT("/tag/edit", controller.UpdateTagHandler)
		v1.DELETE("/tag/delete/:id", controller.DeleteTagHandler)
	}
	return r
}