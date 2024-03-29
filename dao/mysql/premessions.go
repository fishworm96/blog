package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

// 获取一级菜单
func GetMenuList() (menus []*models.MenuDetail, err error) {
	sqlStr := `select id, title, icon, path, type, module_id, is_show from access where type = 1`
	if err = db.Select(&menus, sqlStr); err != nil {
		zap.L().Error("there is no menus in db")
		err = nil
	}
	return
}

// 获取多级菜单列表
func GetChildrenMenuListByMenuId(id int64) (children []*models.MenuDetail, err error) {
	sqlStr := `select id, title, icon, path, type, module_id, is_show from access where module_id = ?`
	if err = db.Select(&children, sqlStr, id); err != nil {
		zap.L().Error("there is no children in db")
		err = nil
		return
	}
	return
}

// 根据菜单id获取单条菜单信息
func GetMenuByMenuId(id int64) (menu *models.MenuDetailInfo, err error) {
	menu = new(models.MenuDetailInfo)
	sqlStr := `select a.id, a.title, a.icon, a.path, a.type, a.module_id, a.is_show, b.title as parent_name
	from access a 
	left join access b on a.module_id = b.id 
	where a.id = ?`
	err = db.Get(menu, sqlStr, id)
	return
}

func GetMenu(id int64) (menu *models.MenuDetail, err error) {
	menu = new(models.MenuDetail)
	sqlStr := `select id, title, icon, path, type, module_id from access where id = ?`
	err = db.Get(menu, sqlStr, id)
	return
}

func GetMenuByUserId(id int64) (menu []*models.MenuDetail, err error) {
	sqlStr := `select id, title, icon, path, type, module_id, is_show from access where id in (
		select access_id from role_access where role_id = (
		select role_id from user where user_id = ?
		) ORDER BY id
	)`
	err = db.Select(&menu, sqlStr, id)
	return
}

// 添加菜单
func AddMenu(m *models.ParamMenu) (err error) {
	sqlStr := `insert into access(title, icon, path, type, module_id) values(?, ?, ?, ?, ?)`
	ret, err := db.Exec(sqlStr, m.Title, m.Icon, m.Path, m.Type, m.ModuleID)
	if err != nil {
		zap.L().Error("add menu failed", zap.Error(err))
		return ErrorAddFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorAddFailed
	}
	return
}

func UpdateMenuById(m *models.ParamUpdateMenu) (err error) {
	sqlStr := `update access set title = ?, icon = ?, path = ?, type = ?, module_id = ?, is_show = ? where id = ?`
	ret, err := db.Exec(sqlStr, m.Title, m.Icon, m.Path, m.Type, m.ModuleID, m.IsShow, m.ID)
	if err != nil {
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorUpdateFailed
	}
	return
}

func UpdateMenuStatus(s models.ParamsMenuStatus) (err error) {
	sqlStr := `
	UPDATE access
	SET is_show = ?
	WHERE id = ?
	`
	_, err = db.Exec(sqlStr, s.IsShow, s.Id)
	return
}

func DeleteMenuById(id int64) (err error) {
	sqlStr := `
	DELETE a, ra 
	FROM access a
	LEFT JOIN role_access ra
	ON a.id = ra.access_id
	WHERE a.id = ?
	`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		zap.L().Error("Delete failed", zap.Error(err))
		return ErrorDeleteFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorMenuNotExist
	}
	return
}
