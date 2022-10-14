package controller

import (
	"blog/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRoleInfoHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	data, err := logic.GetRoleInfoByUserIdHandler(userID)
	if err != nil {
		zap.L().Error("logic.GetRoleInfoByUserIdHandler failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	ResponseSuccess(c, data)
}