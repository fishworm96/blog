package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

// 获取一级菜单
func GetMenuList() (menus []*models.MenuDetail, err error) {
	sqlStr := `select id, title, icon, path, module_id from access where type = 1`
	if err = db.Select(&menus, sqlStr); err != nil {
		zap.L().Error("there is no menus in db")
		err = nil
	}
	return
}

// 获取多级菜单列表
func GetChildrenMenuListByMenuId(id int64) (children []*models.MenuDetail, err error) {
	sqlStr := `select id, title, icon, path, module_id from access where module_id = ?`
	if err = db.Select(&children, sqlStr, id); err != nil {
		zap.L().Error("there is no children in db")
		err = nil
		return
	}
	return
}

// 添加菜单
func AddMenu(m *models.ParamMenu) (err error) {
	sqlStr := `insert into access(title, icon, path, type, module_id) values(?, ?, ?, ?, ?)`
	ret, err := db.Exec(sqlStr, m.Title, m.Icon, m.Path, m.Type, m.ModuleId)
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
	sqlStr := `update access set title = ?, icon = ?, path = ?, type = ?, module_id = ? where id = ?`
	ret, err := db.Exec(sqlStr, m.Title, m.Icon, m.Path, m.Type, m.ModuleId, m.Id)
	if err != nil {
		zap.L().Error("update menu failed", zap.Error(err))
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorUpdateFailed
	}
	return
}