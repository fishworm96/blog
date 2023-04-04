package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"strconv"

	"go.uber.org/zap"
)

// GetCommunityList 获取列表数据
func GetCommunityList() (community []*models.Community, err error) {
	return mysql.GetCommunityList()
}

//  GetCommunityDetail 根据id获取列表数据
func GetCommunityDetail(id, page, size int64) (data *models.CommunityPost, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = new(models.CommunityPost)

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID failed", zap.Int64("post.AuthorID", post.AuthorID), zap.Error(err))
			continue
		}

		ID, err := strconv.ParseInt(post.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		tags, err := mysql.GetTagNameByPostId(ID)
		if err != nil {
			return nil, err
		}
		postDetail := &models.ApiPostDetailList{
			AuthorName:      user.NickName,
			Post:            post,
			Tag:             tags,
		}
		data.ApiPostDetailList = append(data.ApiPostDetailList, postDetail)
	}

	totalPages, err := mysql.GetTotalPages()
	if err != nil {
		zap.L().Error("mysql.GetTotalPages(), failed", zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail(), failed", zap.Error(err))
		return
	}
	
	data.CommunityDetail = community
	data.TotalPages = totalPages
	return
}