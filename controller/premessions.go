package controller

import (
	"blog/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMenuHandler(c *gin.Context) {
	data, err := logic.GetMenuList()
	if err != nil {
		zap.L().Error("logic.GetMenuHandler() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}