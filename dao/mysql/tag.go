package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

func CreateTag(name string) (err error) {
	sqlStr := `insert into tag(tag_name) values (?)`
	_, err = db.Exec(sqlStr, name)
	return err
}

func GetTagList() (tags []*models.Tag, err error) {
	sqlStr := `select id, tag_name from tag`
	if err = db.Select(&tags, sqlStr); err != nil {
		zap.L().Error("there is no tag in db")
		err = nil
	}
	return
}