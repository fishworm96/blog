package controller

import (
	"blog/dao/mysql"
	"blog/logic"
	"blog/models"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func UpdateRoleHandler(c *gin.Context) {
	r := new(models.RoleMenu)
	if err := c.ShouldBindJSON(r); err != nil {
		zap.L().Debug("c.ShouldBindJSON(r) failed", zap.Any("err", err))
		zap.L().Error("update role menu with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	if err := logic.UpdateRoleMenu(r); err != nil {
		zap.L().Error("logic.UpdateRoleMenu(r) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUpdateFailed) {
			ResponseError(c, CodeUpdateFailed)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
