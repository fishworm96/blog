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

// 根绝 RoleId 获取角色
func GetRoleByRoleId(id int64) (role models.Role ,err error) {
	sqlStr := `
	SELECT id, title, description
	FROM role
	WHERE id = ?
	`
	err = db.Get(&role, sqlStr, id)
	return
}

func UpdateRole(role models.ParamsRole) (err error) {
	sqlStr := `
	UPDATE role
	SET title = ?, description = ?
	WHERE id = ?
	`
	ret, err := db.Exec(sqlStr, role.Title, role.Description, role.RoleID)
	if err != nil {
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorTagNotExist
	}
	return err
}

// 根据 id 查询权限
func GetRoleAccessById(id int64) (access_id []string, err error) {
	sqlStr := `
	SELECT access_id
	FROM role_access
	WHERE role_id = ?
	`
	err = db.Select(&access_id, sqlStr, id)
	return
}

// 根据 userID 删除权限
func DeleteRoleAccessByUserID(uid int64) (err error) {
	sqlStr := `delete from role_access where role_id = ?`
	_, err = db.Exec(sqlStr, uid)
	return
}

// 添加权限
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

// 创建角色
func CreateRole(role models.Role) (err error) {
	sqlStr := `
	INSERT INTO role(title, description)
	VALUES (?, ?)
	`
	_, err = db.Exec(sqlStr, role.Title, role.Description)
	return
}

// 通过 role id 删除角色和角色权限
func DeleteRoleAccessByRoleId(id int64) (err error) {
	sqlStr := `
	DELETE r, ra 
	FROM role r 
	LEFT JOIN role_access ra 
	ON r.id = ra.role_id 
	WHERE r.id = ?
	`
	_, err = db.Exec(sqlStr, id)
	return
}