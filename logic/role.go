package logic

import (
	"blog/dao/mysql"
	"blog/models"

	"go.uber.org/zap"
)

// 获取角色列表
func GetRole() (role []*models.Role, err error) {
	return mysql.GetRole()
}

// 根据用户 id 获取信息
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

// 根绝 id 获取角色权限
func GetRoleAccessById(id int64) (data []string, err error) {
	return mysql.GetRoleAccessById(id)
}

func CreateRole(role models.Role) (err error) {
	return mysql.CreateRole(role)
}

// 修改角色权限
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

// 删除角色
func DeleteRoleById(id int64) (err error) {
	return mysql.DeleteRoleAccessByRoleId(id)
}