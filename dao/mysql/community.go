package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		zap.L().Error("there is no community in db")
		err = nil
	}
	return
}