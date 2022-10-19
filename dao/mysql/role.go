package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

func GetRoleHandler(uid int64) (role *models.Role, err error) {
	role = new(models.Role)
	sqlStr := `select role_id, title, description from role where role_id = (
		select role_id from user where user_id = ?
	)`
	err = db.Get(role, sqlStr, uid)
	return
}

func DeleteRoleAccessByUserID(uid int64) (err error) {
	sqlStr := `delete from role_access where role_id = ?`
	ret, err := db.Exec(sqlStr, uid)
	if err != nil {
		zap.L().Error("delete role access failed", zap.Error(err))
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorUpdateFailed
	}
	return
}

func AddRoleAccess(RoleID, AccessID int64) (err error) {
	sqlStr := `insert into role_access(role_id, access_id) values(?, ?)`
	ret, err := db.Exec(sqlStr, RoleID, AccessID)
	if err != nil {
		zap.L().Error("add role access failed", zap.Error(err))
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorUpdateFailed
	}
	return
}
