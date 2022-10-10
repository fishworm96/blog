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

func AddMenu(m *models.Menu) error {
	return mysql.AddMenu(m)
}