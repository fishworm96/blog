package controller

import (
	"blog/logic"
	"blog/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// CreateTagHandler 创建标签接口
func CreateTagHandler(c *gin.Context) {
	tag := new(models.Tag)
	if err := c.ShouldBindJSON(tag); err != nil {
		zap.L().Error("create post with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	
	if err := logic.CreateTag(tag); err != nil {
		zap.L().Error("logic.CreateTag(name) failed", zap.Error(err))
		ResponseError(c, COdeTagExist)
		return
	}

	ResponseSuccess(c, nil)
}

// GetTagHandler 获取标签列表接口
func GetTagListHandler(c *gin.Context) {
	data, err := logic.GetTagList()
	if err != nil {
		zap.L().Error("logic.GetTag() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}