package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

// GetCommunityList 获取列表数据
func GetCommunityList() (community []*models.Community, err error) {
	return mysql.GetCommunityList()
}

//  GetCommunityDetail 根据id获取列表数据
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(id)
}