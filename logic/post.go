package logic

import (
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	"blog/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 生成post id
	p.ID = snowflake.GenID()
	// 保存到数据库
	for _, tagId := range p.Tag {
		err = mysql.CreateTagByPostId(p.ID, tagId)
		if err != nil {
			return err
		}
	}
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}

func AddPostTag(postTag *models.ParamPostAndTag) error {
	return mysql.AddPostTag(postTag)
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
	tags, err := mysql.GetTagsByPostId(post.ID)
	if err != nil {
		return nil, err
	}
	// 接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName:      user.NickName,
		Post:            post,
		CommunityDetail: community,
		Tag: tags,
	}
	return
}

func GetPostList(page int64, size int64) (data []*models.ApiPostList, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostList, 0, len(posts))

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
		data = append(data, postDetail)
	}
	return
}

func UpdatePost(p *models.ParamPost) (err error) {
	err = mysql.UpdatePost(p)
	if err != nil {
		return err
	}
	err = mysql.DeleteTagByPostID(p.PostID)
	for _, tag := range p.Tag {
		err = mysql.CreateTagByPostId(p.PostID, tag)
		if err != nil {
			return
		}
	}
	return
}

func DeletePostById(pid int64) error {
	return mysql.DeletePostById(pid)
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostList, err error) {
	// 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Error("redis.GetPostIDsOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 根据id去mysql数据库查询帖子详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据社区id查询详细信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostList{
			AuthorName:      user.NickName,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostList, err error) {
	// 去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetCommunityPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("redis.GetCommunityPostIDsInOrder(p)", zap.Any("ids", ids))
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetCommunityPostList", zap.Any("posts", posts))
	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostList{
			AuthorName:      user.NickName,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostList, err error) {
	// 根据请求参数的不同，执行不同的逻辑
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetCommunityListNew failed", zap.Error(err))
		return nil, err
	}
	return
}
