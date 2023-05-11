package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"strconv"

	"go.uber.org/zap"
)

func CreateTag(tag *models.Tag) error {
	// 判断标签是否存在
	if err := mysql.CheckTagExist(tag.Name); err != nil {
		return err
	}
	return mysql.CreateTag(tag.Name)
}

func GetTagByName(name string, page, size int64) (data *models.ApiTagDetail, err error) {
	// 获取标签名称
	tag, err := mysql.GetTagByName(name)
	if err != nil {
		zap.L().Error("mysql.GetTagByName(name) failed", zap.Error(err))
		return
	}

	// 通过标签名称获取帖子信息
	posts, err := mysql.GetPostByTagId(tag.ID, page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostByTagId(tag.ID) failed", zap.Int64("tag.ID", tag.ID), zap.Error(err))
	}

	postList := make([]*models.ApiPostDetailList, 0, len(posts))

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
			CommunityDetail: community,
			Tag:             tags,
		}
		postList = append(postList, postDetail)
	}
	data = &models.ApiTagDetail{
		Id:   tag.ID,
		Name: tag.Name,
		Post: postList,
	}
	totalTag, err := mysql.GetTotalByTag(name)
	if err != nil {
		zap.L().Error("mysql.GetTotalByTag(name) failed", zap.String("tagName", name))
		return
	}
	data.TotalPages = totalTag

	return
}

func GetTagList() (data []*models.Tag, err error) {
	return mysql.GetTagList()
}

func UpdateTag(tag *models.Tag) (err error) {
	err = mysql.UpdateTag(tag)
	if err != nil {
		zap.L().Error("mysql.UpdateTag(tag)", zap.Error(err))
		return
	}
	return
}

func DeleteTagById(tid int64) error {
	return mysql.DeleteTagById(tid)
}
