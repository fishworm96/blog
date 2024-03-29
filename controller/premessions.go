package controller

import (
	"blog/logic"
	"blog/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 获取菜单列表
func GetMenuHandler(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if ok {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			zap.L().Error("get menu detail with invalid param", zap.Int64("id", id), zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		data, err := logic.GetMenuByMenuId(id)
		if err != nil {
			zap.L().Error("logic.getMenuByMenuId failed", zap.Error(err))
			ResponseError(c, CodeMenuNotExist)
			return
		}
		ResponseSuccess(c, data)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID(c) failed", zap.Any("userID", userID))
		ResponseError(c, CodeNeedLogin)
		return
	}
	data, err := logic.GetMenuByUserId(userID)
	if err != nil {
		zap.L().Error("logic.GetMenuByUserId(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// func GetMenuByMenuIdHandler(c *gin.Context) {
// 	idStr := c.Query("id")
// 	zap.L().Error("id", zap.Any("id", idStr))
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		zap.L().Error("get menu detail with invalid param", zap.Int64("id", id), zap.Error(err))
// 		ResponseError(c, CodeInvalidParam)
// 		return
// 	}
// 	data, err := logic.GetMenuByMenuId(id)
// 	if err != nil {
// 		zap.L().Error("logic.getMenuByMenuId failed", zap.Error(err))
// 		ResponseError(c, CodeServerBusy)
// 		return
// 	}

// 	ResponseSuccess(c, nil)
// }

// 添加菜单
func AddMenuHandler(c *gin.Context) {
	m := new(models.ParamMenu)
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

// 修改菜单
func UpdateMenuHandler(c *gin.Context) {
	m := new(models.ParamUpdateMenu)
	if err := c.ShouldBindJSON(m); err != nil {
		zap.L().Debug("c.ShouldBindJSON(m) failed", zap.Any("err", err), zap.Any("params", m))
		zap.L().Error("update menu with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	err := logic.UpdateMenu(m)
	if err != nil {
		zap.L().Error("logic.UpdateMenu(m) failed", zap.Error(err))
		ResponseError(c, CodeUpdateFailed)
		return
	}

	ResponseSuccess(c, nil)
}

func UpdateStatusHandler(c *gin.Context) {
	var status models.ParamsMenuStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		zap.L().Error("c.ShouldBindJSON(&status) failed", zap.Any("status", status), zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	err := logic.UpdateMenuStatus(status)
	if err != nil {
		zap.L().Error("logic.UpdateMenuStatus(status) failed", zap.Error(err))
		ResponseError(c, CodeUpdateFailed)
		return
	}

	ResponseSuccess(c, nil)
}

// 删除菜单
func DeleteMenuHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("delete post detail with invalid param", zap.Int64("id", id), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	state, err := logic.DeleteMenu(id)
	if err != nil {
		zap.L().Error("logic.DeleteMenu(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if !state {
		ResponseError(c, CodeDeleteFailed)
		return
	}

	ResponseSuccess(c, nil)
}
