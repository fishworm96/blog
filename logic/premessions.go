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
		childrenList, err := mysql.GetChildrenMenuListByMenuId(menu.ID)
		if err != nil {
			zap.L().Error("mysql.GetChildrenMenuListByMenuId(menu.Id) failed", zap.Int64("menu.Id", menu.ID), zap.Error(err))
			continue
		}
		for _, child := range childrenList {
			chil, err := mysql.GetChildrenMenuListByMenuId(child.ID)
			if err != nil {
				zap.L().Error("mysql.GetChildrenMenuListByMenuId(menu.Id) failed", zap.Int64("menu.Id", menu.ID), zap.Error(err))
				continue
			}
			child.Children = chil
		}
		menu.Children = childrenList
		data = append(data, menu)
	}
	return
}

func getTreeRecursive(list []*models.MenuDetail, parentId int64) []*models.MenuDetail {
	res := make([]*models.MenuDetail, 0)
	for _, v := range list {
		if v.ModuleID == parentId {
			v.Children = getTreeRecursive(list, v.ID)
			res = append(res, v)
		}
	}
	return res
}

func GetMenuByUserId(id int64) (data []*models.MenuDetail, err error) {
	info, err := mysql.GetUserInfoByUserId(id)
	if err != nil {
		zap.L().Error("mysql.GetUserInfoByUserId(id) failed", zap.Int64("id", id), zap.Error(err))
		return
	}
	if info.IsSuper == 1 {
		data, err = GetMenuList()
		return
	}
	menus, err := mysql.GetMenuByUserId(id)
	data = getTreeRecursive(menus, 0)
	return
}

func GetMenuByMenuId(id int64) (data *models.MenuDetailInfo, err error) {
	data, err = mysql.GetMenuByMenuId(id)
	return
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
	menu, err := mysql.GetChildrenMenuListByMenuId(m.ID)
	if err != nil {
		zap.L().Error("mysql.GetMenu(id) failed", zap.Int64("id", id), zap.Error(err))
		return
	}
	if menu != nil {
		return false, err
	}
	return true, mysql.DeleteMenuById(id)
}
