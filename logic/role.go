package logic

import (
	"blog/dao/mysql"
	"blog/models"

	"go.uber.org/zap"
)

func GetRole() (role []*models.Role, err error) {
	return mysql.GetRole()
}

func GetRoleInfoByUserIdHandler(uid int64) (data *models.RoleInfo, err error) {
	user, err := mysql.GetUserById(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
		return
	}
	role, err := mysql.GetRoleById(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
		return
	}
	data = &models.RoleInfo{
		Username:    user.NickName,
		Title:       role.Title,
		Description: role.Description,
	}
	return
}

func UpdateRoleMenu(role *models.RoleMenu) (err error) {
	if err = mysql.DeleteRoleAccessByUserID(role.RoleID); err != nil {
		return
	}
	for _, r := range role.AccessID {
		err = mysql.AddRoleAccess(role.RoleID, r)
		continue
	}
	return
}
