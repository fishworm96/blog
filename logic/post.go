package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	// 生成post id
	p.ID = snowflake.GenID()
	// 保存到数据库
	return mysql.CreatePost(p)
}

// GetPostById 根据帖子id查询帖子详情数据
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想要的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid)", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("post.AuthorID", post.AuthorID), zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID) failed", zap.Int64("post.communityID", post.CommunityID), zap.Error(err))
		return
	}
	// 接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName: user.Username,
		Post: post,
		CommunityDetail: community,
	}
	return
}