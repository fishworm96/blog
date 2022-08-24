package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func GetCommunityList() (community []*models.Community, err error) {
	return mysql.GetCommunityList()
}