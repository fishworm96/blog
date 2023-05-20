package middlewares

import (
	"blog/controller"
	"blog/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

const Authorization = "Authorization"

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(Authorization)
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分隔
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1]是获取到tokenString, 我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		newToken, err, ok := jwt.RefreshToken(parts[1], mc.Username, mc.UserID )
		// 返回请求头 token
		if ok {
			c.Header(Authorization, newToken)
		}
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理请求的函数中，可以用c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
