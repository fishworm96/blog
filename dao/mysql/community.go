package mysql

import (
	"blog/models"
	"database/sql"

	"go.uber.org/zap"
)

// GetCommunityList 获取社区列表
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `
	SELECT id, name, image
	FROM community
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
	sqlStr := `
	SELECT id, name, image, introduction, create_time, update_time
	FROM community WHERE id = ?
	`
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
		INSERT INTO community(name, introduction, image)
		VALUES(?, ?, ?)
	`
	_, err = db.Exec(sqlStr, c.Name, c.Introduction, c.Image)
	return
}

func UpdateCommunity(c models.Community) (err error) {
	sqlStr := `
	UPDATE community
	SET name = ?, image = ?
	WHERE id = ?
	`
	ret, err := db.Exec(sqlStr, c.Name, c.Image, c.ID)
	if err != nil {
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorCommunityNotExist
	}
	return
}

func DeleteCommunity(id int64) (err error) {
	sqlStr := `
	DELETE FROM community WHERE id = ?
	`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		return ErrorDeleteFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorCommunityNotExist
	}
	if err != nil {
		return err
	}
	return
}