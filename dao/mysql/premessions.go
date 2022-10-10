package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

func GetMenuList() (menus []*models.MenuDetail, err error) {
	sqlStr := `select id, title, icon, path, module_id from access where type = 1`
	if err = db.Select(&menus, sqlStr); err != nil {
		zap.L().Error("there is no menus in db")
		err = nil
	}
	return
}

func GetChildrenMenuListByMenuId(id int64) (children []*models.MenuDetail, err error) {
	sqlStr := `select id, title, icon, path, module_id from access where module_id = ?`
	if err = db.Select(&children, sqlStr, id); err != nil {
		zap.L().Error("there is no children in db")
		err = nil
		return
	}
	return
}

func AddMenu(m *models.Menu) (err error) {
	sqlStr := `
		insert into access(title, icon, path, type, module_id) values(?, ?, ?, ?, ?)
	`
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
