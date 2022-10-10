package controller

import (
	"blog/logic"
	"blog/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func AddMenuHandler(c *gin.Context) {
	m := new(models.Menu)
	if err := c.ShouldBindJSON(m); err != nil {
		zap.L().Debug("c.ShouldBindJSON(m) failed", zap.Any("err", err))
		zap.L().Error("add menu with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	err := logic.AddMenu(m)
	if err != nil {
		zap.L().Error("logic.AddMenuHandler(m) failed", zap.Error(err))
		ResponseError(c, CodeAddFailed)
		return
	}

	ResponseSuccess(c, nil)
}