package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

// CheckTagExist 检查指定标签是否存在
func CheckTagExist(tagName string) (err error) {
	sqlStr := `select count(id) from tag where tag_name = ?`
	var count int
	if err := db.Get(&count, sqlStr, tagName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorTagExist
	}
	return
}

// CreateTag 向数据库插入一条新的标签
func CreateTag(name string) (err error) {
	sqlStr := `insert into tag(tag_name) values (?)`
	_, err = db.Exec(sqlStr, name)
	return err
}

// GetTagList 向数据库查询标签列表
func GetTagList() (tags []*models.Tag, err error) {
	sqlStr := `select id, tag_name from tag`
	if err = db.Select(&tags, sqlStr); err != nil {
		zap.L().Error("there is no tag in db")
		err = nil
	}
	return
}

// UpdateTag 更新数据库标签名称
func UpdateTag(tag *models.Tag) (err error) {
	sqlStr := `update tag set tag_name = ? where id = ?`
	_, err = db.Exec(sqlStr, tag.Name, tag.Id)
	return err
}