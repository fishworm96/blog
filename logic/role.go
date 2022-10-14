package logic

import (
	"blog/dao/mysql"
	"blog/models"

	"go.uber.org/zap"
)

func GetRoleInfoByUserIdHandler(uid int64) (data *models.UserInfo, err error) {
	user, err := mysql.GetUserById(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
	}
	role, err := mysql.GetRoleHandler(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
	}
	data = &models.UserInfo{
		Username:    user.Username,
		Title:       role.Title,
		Description: role.Description,
	}
	return
}