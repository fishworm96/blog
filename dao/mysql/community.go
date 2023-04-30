package mysql

import (
	"blog/models"
	"database/sql"

	"go.uber.org/zap"
)

// GetCommunityList 获取社区列表
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `
	SELECT c.id, c.name, i.image_url
	FROM community c
	INNER JOIN image i ON c.image_md5 = i.md5;
	`
	if err := db.Select(&communityList, sqlStr); err != nil {
		zap.L().Error("there is no community in db")
		err = nil
	}
	return
}

// GetCommunityDetail 获取社区详细信息
func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select id, name, introduction, create_time, update_time from community where id = ?`
	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}

// GetCommunityDetailById 根据社区id获取社区信息
func GetCommunityDetailById(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select id, name, introduction, create_time, update_time from community where id = ?`
	err = db.Get(community, sqlStr, id)
	return
}

func GetTotalCategory() (totalCategory int64, err error) {
	sqlStr := `select count(id) from community`
	err = db.Get(&totalCategory, sqlStr)
	return
}

func CreateCommunity(c models.CommunityCreateDetail) (err error) {
	sqlStr := `
		INSERT INTO community(name, introduction, image_md5)
		VALUES(?, ?, ?)
	`
	_, err = db.Exec(sqlStr, c.Name, c.Introduction, c.Md5)
	return
}