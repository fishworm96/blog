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

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 根据id获取社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取分页参数
	page, size := getPageInfo(c)

	// 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id, page, size)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// 创建分类
func CreateCommunity(c *gin.Context) {
	var community models.CommunityCreateDetail
	if err := c.ShouldBindJSON(&community); err != nil {
		zap.L().Debug("c.ShouldBind(&community) err", zap.Any("err", err))
		zap.L().Error("create community with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate((trans))))
		return
	}

	err := logic.CreateCommunity(community)
	if err != nil {
		zap.L().Error("logic.CreateCommunity(community) failed", zap.Error(err))
		ResponseError(c, CodeAddFailed)
		return
	}

	ResponseSuccess(c, nil)
}

func UpdateCommunity(c *gin.Context) {
	var community models.Community
	if err := c.ShouldBindJSON(&community); err != nil {
		zap.L().Error("update community with param", zap.Any("community", community), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.UpdateCommunity(community); err != nil {
		zap.L().Error("logic.UpdateCommunity(community) failed", zap.Any("community", community), zap.Error(err))
		ResponseError(c, CodeUpdateFailed)
		return
	}

	ResponseSuccess(c, nil)
}

func DeleteCommunity(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("delete community with invalid param", zap.Int64("id", id), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.DeleteCommunity(id); err != nil {
		zap.L().Error("logic.DeleteCommunity(id) failed", zap.Int64("id", id), zap.Error(err))
		if errors.Is(err, mysql.ErrorCommunityNotExist) {
			ResponseError(c, CodeCommunityNotExist)
			return
		}
		ResponseError(c, CodeDeleteFailed)
		return
	}

	ResponseSuccess(c, nil)
}