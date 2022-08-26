package logic

import (
	"blog/dao/mysql"
	"blog/dao/redis"
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
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(page int64, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID failed", zap.Int64("post.AuthorID", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID), failed", zap.Int64("post.CommunityID", post.CommunityID))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func UpdatePost(pid int64, p *models.ParamPost) error {
	return mysql.UpdatePost(pid, p)
}

func DeletePostById(pid int64) error {
	return mysql.DeletePostById(pid)
}

func CreAtePost(p *models.Post) (err error) {
	// 生成post id
	p.ID = snowflake.GenID()
	// 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return
}