package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateTag(tag *models.Tag) error {
	return mysql.CreateTag(tag.Name)
}

func GetTagList() (data []*models.Tag, err error) {
	return mysql.GetTagList()
}