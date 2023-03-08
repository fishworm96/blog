package logic

import (
	"blog/dao/mysql"
	"blog/models"

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
	posts, err := mysql.GetPostByTagName(tag.Name, page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostByTagName(tag.Name) failed", zap.String("tag.Name", tag.Name), zap.Error(err))
	}

	postList := make([]*models.ApiPostList, 0, len(posts))

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
		tags, err := mysql.GetTagNameByPostId(post.ID)
		if err != nil {
			return nil, err
		}
		postDetail := &models.ApiPostList{
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
	// data = new(models.ApiTagDetail)
	// data.Id = tag.Id
	// data.Name = tag.Name
	// data.Post = postList
	return
}

func GetTagList() (data []*models.Tag, err error) {
	return mysql.GetTagList()
}

func UpdateTag(tag *models.Tag) error {
	// 判断标签是否存在
	if err := mysql.CheckTagExist(tag.Name); err != nil {
		return err
	}
	return mysql.UpdateTag(tag)
}

func DeleteTagById(tid int64) error {
	return mysql.DeleteTagById(tid)
}
