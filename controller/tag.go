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
		ResponseError(c, CodeTagExist)
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

// GetTagDetailHandler 根据标签name获取标签信息接口
func GetTagDetailHandler(c *gin.Context) {
	name := c.Param("name")
	page, size := getPageInfo(c)
	data, err := logic.GetTagByName(name, page, size)
	if err != nil {
		zap.L().Error("logic.GetTagById(tid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// UpdateTagHandler 修改标签
func UpdateTagHandler(c *gin.Context) {
	tag := new(models.Tag)
	if err := c.ShouldBindJSON(tag); err != nil {
		zap.L().Error("update tag with param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.UpdateTag(tag); err != nil {
		zap.L().Error("logic.UpdateTag(tag) failed", zap.Any("tid", tag), zap.Error(err))
		if errors.Is(err, mysql.ErrorUpdateFailed) {
			ResponseError(c, CodeUpdateFailed)
			return
		}
		ResponseError(c, CodeUpdateFailed)
		return
	}
	ResponseSuccess(c, nil)
}

// DeleteTagHandler 删除标签接口
func DeleteTagHandler(c *gin.Context) {
	tidStr := c.Param("ids")
	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		zap.L().Error("delete tag with invalid param", zap.Int64("tid", tid), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id删除标签
	if err := logic.DeleteTagById(tid); err != nil {
		zap.L().Error("logic.DeleteTagById(tid) failed", zap.Int64("tid", tid), zap.Error(err))
		if errors.Is(err, mysql.ErrorTagNotExist) {
			ResponseError(c, CodeTagNotExist)
			return
		}
		ResponseError(c, CodeDeleteFailed)
		return
	}
	ResponseSuccess(c, nil)
}