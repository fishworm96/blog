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