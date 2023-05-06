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

// 获取用户权限
func GetUserRoleHandler(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	// 获取单条数据
	if ok {
		zap.L().Error("get role with param", zap.String("idStr", idStr))
		ResponseError(c, CodeInvalidParam)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
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

// 获取权限列表
func GetRoleHandler(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	// 获取单条数据
	if ok {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			zap.L().Error("get role with param", zap.Int64("id", id), zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}

		data, err := logic.GetRoleAccessById(id)
		if err != nil {
			zap.L().Error("logic.GetRoleAccessById(id) failed", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}

		ResponseSuccess(c, data)
		return
	}

	// 获取权限列表
	data, err := logic.GetRole()
	if err != nil {
		zap.L().Error("logic.GetRole() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// 创建角色
func CreateRoleHandler(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		zap.L().Error("create role with parma", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.CreateRole(role); err != nil {
		zap.L().Error("logic.CreateRole(role) failed", zap.Error(err))
		ResponseError(c, CodeAddFailed)
		return
	}

	ResponseSuccess(c, nil)
}

// 删除角色
func DeleteRoleAccessByRoleIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("delete role with invalid param", zap.Int64("id", id), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.DeleteRoleById(id); err != nil {
		zap.L().Error("logic.DeleteRoleById(id) failed", zap.Error(err))
		ResponseError(c, CodeDeleteFailed)
		return
	}

	ResponseSuccess(c, nil)
}

func UpdateRoleHandler(c *gin.Context) {
	var role models.ParamsRole
	if err := c.ShouldBindJSON(&role); err != nil {
		zap.L().Error("update role with invalid param", zap.Any("role", role), zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	if err := logic.UpdateRole(role); err != nil {
		zap.L().Error("logic.UpdateRole(role) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUpdateFailed) {
			ResponseError(c, CodeUpdateFailed)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// 修改角色权限
func UpdateRoleAndAccessHandler(c *gin.Context) {
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
		zap.L().Error("logic.UpdateRoleMenu(r) failed", zap.Any("r", r), zap.Error(err))
		if errors.Is(err, mysql.ErrorUpdateFailed) {
			ResponseError(c, CodeUpdateFailed)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
