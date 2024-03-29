package routes

import (
	"blog/controller"
	"blog/logger"
	"blog/middlewares"
	"blog/websocket"

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

	hub := websocket.NewHub()
	go hub.Run()

	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.Cors())

	r.Static("/static", "./static")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")

	v1.POST("/email", controller.GetEmailCode)
	v1.POST("/email_login", controller.EmailLoginHandler)
	v1.POST("/signUp", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/post", controller.GetPostListHandler)
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.GET("/tag", controller.GetTagListHandler)
	v1.GET("/tag/:name", controller.GetTagDetailHandler)
	v1.GET("/search", controller.SearchArticles)
	v1.GET("/ws", func(c *gin.Context) {
		websocket.HttpController(c, hub)
	})

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/info/:id", controller.GetUserInfoHandler)
		v1.GET("/info_list", controller.GetUserInfoListHandler)
		v1.POST("/upload_avatar", controller.UploadAvatar)
		v1.POST("/upload_image", controller.UploadImage)
	}
	{
		v1.GET("/user_role", controller.GetUserRoleHandler)
		v1.GET("/role", controller.GetRoleHandler)
		v1.POST("/role", controller.CreateRoleHandler)
		v1.PUT("/role_access", controller.UpdateRoleAndAccessHandler)
		v1.PUT("/role", controller.UpdateRoleHandler)
		v1.DELETE("/role/:id", controller.DeleteRoleAccessByRoleIdHandler)
	}
	{
		v1.GET("/menu", controller.GetMenuHandler)
		v1.POST("/menu", controller.AddMenuHandler)
		v1.PUT("/menu", controller.UpdateMenuHandler)
		v1.PUT("/menu/status", controller.UpdateStatusHandler)
		v1.DELETE("/menu/:id", controller.DeleteMenuHandler)
	}
	{
		v1.POST("/post", controller.CreatePostHandler)
		v1.POST("/post/add_post_tag", controller.AddPostTagHandler)
		v1.PUT("/post/edit", controller.UpdatePostHandler)
		v1.DELETE("/post/:id", controller.DeletePostHandler)
		v1.POST("/vote", controller.PostVoteController)
	}
	{
		v1.POST("/tag", controller.CreateTagHandler)
		v1.PUT("/tag/edit", controller.UpdateTagHandler)
		v1.DELETE("/tag/:id", controller.DeleteTagHandler)
	}
	{
		v1.POST("/community", controller.CreateCommunity)
		v1.PUT("/community", controller.UpdateCommunity)
		v1.DELETE("/community/:id", controller.DeleteCommunity)
	}
	return r
}
