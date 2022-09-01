package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateTag(tag *models.Tag) error {
	return mysql.CreateTag(tag.Name)
}