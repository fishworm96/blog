package mysql

import (
	"blog/models"
	"database/sql"

	"go.uber.org/zap"
)

// GetCommunityList 获取社区列表
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		zap.L().Error("there is no community in db")
		err = nil
	}
	return
}

// GetCommunityDetail 获取社区详细信息
func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time, update_time from community where community_id = ?`
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
	sqlStr := `select community_id, community_name, introduction, create_time, update_time from community where community_id = ?`
	err = db.Get(community, sqlStr, id)
	return
}

func GetTotalCategory() (totalCategory int64, err error) {
	sqlStr := `select count(id) from community`
	err = db.Get(&totalCategory, sqlStr)
	return
}