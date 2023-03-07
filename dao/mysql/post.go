package mysql

import (
	"blog/models"
	"strings"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, description, content, author_id, community_id) values (?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Description, p.Content, p.AuthorID, p.CommunityID)
	return
}

func AddPostTag(postTag *models.ParamPostAndTag) (err error) {
	sqlStr := `
	insert into post_tag(post_id, tag_name)
	SELECT ?, ?
	from DUAL
	where EXISTS (
		select post.post_id, tag.tag_name 
		from post, tag 
		where post.post_id = ? 
		and tag.tag_name = ?
		)
		`
	ret, err := db.Exec(sqlStr, postTag.ID, postTag.Name, postTag.ID, postTag.Name)
	if err != nil {
		zap.L().Error("add post tag failed", zap.Error(err))
		return ErrorAddFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorAddFailed
	}
	return
}

// GetPostById 根据帖子id获取信息
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	if err != nil {
		return nil, ErrorPostNotExist
	}
	return
}

// GetPostList 查询所有帖子信息
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post order by create_time desc limit ?, ?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page - 1) * size, size)
	return
}

// UpdatePost 更新文章信息
func UpdatePost(p *models.ParamPost) (err error) {
	sqlStr := `update post set title = ?, content = ? where post_id = ?`
	ret, err := db.Exec(sqlStr, p.Title, p.Content, p.PostID)
	if err != nil {
		zap.L().Error("Update failed", zap.Error(err))
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorPostNotExist
	}
	return
}

func DeletePostById(pid int64) (err error) {
	sqlStr := `delete from post where post_id = ?`
	ret, err := db.Exec(sqlStr, pid)
	if err != nil {
		zap.L().Error("Delete failed", zap.Error(err))
		return ErrorDeleteFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorPostNotExist
	}
	return
}

func DeleteTagByPostID(postID int64) (err error) {
	sqlStr := `delete from post_tag where post_id = ?`
	ret, err := db.Exec(sqlStr, postID)
	if err != nil {
		zap.L().Error("Delete failed", zap.Error(err))
		return
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorTagNotExist
	}
	return
}

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id in (?) order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}