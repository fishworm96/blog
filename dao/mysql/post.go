package mysql

import (
	"blog/models"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据帖子id获取信息
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
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
func UpdatePost(pid int64, p *models.ParamPost) (err error) {
	sqlStr := `update post set title = ?, content = ? where post_id = ?`
	_, err = db.Exec(sqlStr, p.Title, p.Content, pid)
	return
}

func DeletePostById(pid int64) (err error) {
	sqlStr := `delete from post where post_id = ?`
	_, err = db.Exec(sqlStr, pid)
	return
}