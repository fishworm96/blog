package logic

import (
	"blog/dao/mysql"
	"blog/models"

	"go.uber.org/zap"
)

func GetMenuList() (data []*models.MenuDetail, err error) {
	menus, err := mysql.GetMenuList()
	if err != nil {
		return
	}
	for _, menu := range menus {
		childrenList, err := mysql.GetChildrenMenuListByMenuId(menu.Id)
		if err != nil {
			zap.L().Error("mysql.GetChildrenMenuListByMenuId(menu.Id) failed", zap.Int64("menu.Id", menu.Id), zap.Error(err))
			continue
		}
		for _, child := range childrenList {
			chil, err := mysql.GetChildrenMenuListByMenuId(child.Id)
			if err != nil {
				zap.L().Error("mysql.GetChildrenMenuListByMenuId(menu.Id) failed", zap.Int64("menu.Id", menu.Id), zap.Error(err))
				continue
			}
			child.Children = chil
		}
		menu.Children = childrenList
		data = append(data, menu)
	}
	return
}

func GetMenu(id int64) (data *models.MenuDetail, err error) {
	return mysql.GetMenu(id)
}

func AddMenu(m *models.ParamMenu) error {
	return mysql.AddMenu(m)
}

func UpdateMenu(m *models.ParamUpdateMenu) error {
	return mysql.UpdateMenuById(m)
}

func DeleteMenu(id int64) (state bool, err error) {
	m, err := mysql.GetMenu(id)
	if err != nil {
		zap.L().Error("mysql.GetMenu(id) failed", zap.Int64("id", id), zap.Error(err))
		return
	}
	menu, err := mysql.GetChildrenMenuListByMenuId(m.Id)
	if err != nil {
		zap.L().Error("mysql.GetMenu(id) failed", zap.Int64("id", id), zap.Error(err))
		return
	}
	if menu != nil {
		return false, err
	}
	return true, mysql.DeleteMenuById(id)
}