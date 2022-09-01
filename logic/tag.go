package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CreateTag(tag *models.Tag) error {
	// 判断标签是否存在
	if err := mysql.CheckTagExist(tag.Name); err != nil {
		return err
	}
	return mysql.CreateTag(tag.Name)
}

func GetTagList() (data []*models.Tag, err error) {
	return mysql.GetTagList()
}

func UpdateTag(tag *models.Tag) error {
	// 判断标签是否存在
	if err := mysql.CheckTagExist(tag.Name); err != nil {
		return err
	}
	return mysql.UpdateTag(tag)
}

func DeleteTagById(tid int64) error {
	return mysql.DeleteTagById(tid)
}
