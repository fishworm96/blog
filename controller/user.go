package controller

import (
	"blog/logic"
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数有误，直接返回错误信息
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationError类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
		// 2.业务逻辑
		if err := logic.SignUp(p); err != nil {
			zap.L().Error("logic.SingUp failed", zap.Error(err))
			c.JSON(http.StatusOK, gin.H{
				"msg": "注册失败",
			})
			return
		}
		// 3.返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "succes",
		})
}