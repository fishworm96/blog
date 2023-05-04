package mysql

import (
	"blog/models"

)

func GetRole() (role []*models.Role, err error) {
	sqlStr := `
	SELECT id, title, description
	FROM role
	`
	db.Select(&role, sqlStr)
	return
}

func GetRoleById(uid int64) (role *models.Role, err error) {
	role = new(models.Role)
	sqlStr := `select id, title, description from role where id = (
		select role_id from user where user_id = ?
	)`
	err = db.Get(role, sqlStr, uid)
	return
}

func DeleteRoleAccessByUserID(uid int64) (err error) {
	sqlStr := `delete from role_access where id = ?`
	ret, err := db.Exec(sqlStr, uid)
	if err != nil {
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
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorUpdateFailed
	}
	return
}
