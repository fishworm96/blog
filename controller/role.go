package controller

import (
	"blog/dao/mysql"
	"blog/logic"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRoleInfoHandler(c *gin.Context) {
	pidStr := c.Param("id")
	id, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get role with param", zap.Int64("id", id), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetRoleInfoByUserIdHandler(id)
	if err != nil {
		zap.L().Error("logic.GetRoleInfoByUserIdHandler failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}