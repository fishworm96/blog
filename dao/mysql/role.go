package mysql

import "blog/models"

func GetRoleHandler(uid int64) (role *models.Role, err error) {
	role = new(models.Role)
	sqlStr := `select role_id, title, description from role where role_id = (
		select role_id from user where user_id = ?
	)`
	err = db.Get(role, sqlStr, uid)
	return
}