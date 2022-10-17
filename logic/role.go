package logic

import (
	"blog/dao/mysql"
	"blog/models"

	"go.uber.org/zap"
)

func GetRoleInfoByUserIdHandler(uid int64) (data *models.RoleInfo, err error) {
	user, err := mysql.GetUserById(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
		return
	}
	role, err := mysql.GetRoleHandler(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
		return
	}
	data = &models.RoleInfo{
		Username:    user.Username,
		Title:       role.Title,
		Description: role.Description,
	}
	return
}