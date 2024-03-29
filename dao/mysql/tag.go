package mysql

import (
	"blog/models"

	"go.uber.org/zap"
)

// CheckTagExist 检查指定标签是否存在
func CheckTagExist(tagName string) (err error) {
	sqlStr := `select count(id) from tag where tag_name = ?`
	var count int
	if err := db.Get(&count, sqlStr, tagName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorTagExist
	}
	return
}

// CreateTag 向数据库插入一条新的标签
func CreateTag(name string) (err error) {
	sqlStr := `insert into tag(tag_name) values (?)`
	_, err = db.Exec(sqlStr, name)
	return
}

func CreateTagByPostId(postID, tagID int64) (err error) {
	sqlStr := `insert into post_tag(post_id, tag_id) values (?, ?)`
	_, err = db.Exec(sqlStr, postID, tagID)
	return
}

// GetTagList 向数据库查询标签列表
func GetTagList() (tags []*models.Tag, err error) {
	sqlStr := `select id, tag_name from tag`
	if err = db.Select(&tags, sqlStr); err != nil {
		zap.L().Error("there is no tag in db")
		err = nil
	}
	return
}

func GetTagByName(name string) (tag *models.Tag, err error) {
	tag = new(models.Tag)
	sqlStr := `select id, tag_name from tag where tag_name = ?`
	if err = db.Get(tag, sqlStr, name); err != nil {
		zap.L().Error("there is no tag in db", zap.String("name", name), zap.Error(err))
		err = nil
	}
	return
}

func GetTagNameByPostId(pID int64) (tags []string, err error) {
	sqlStr := `
	select distinct tag.tag_name
	FROM post
	INNER JOIN post_tag on ? = post_tag.post_id
	INNER JOIN tag on post_tag.tag_id = tag.id
	`
	if err = db.Select(&tags, sqlStr, pID); err != nil {
		zap.L().Error("there is no tag in db", zap.Int64("pid", pID), zap.Error(err))
		err = nil
	}
	return
}

func GetTagsByPostId(pID int64) (tags []*models.Tag, err error) {
	sqlStr := `
	select distinct tag.id, tag.tag_name
	FROM post
	INNER JOIN post_tag on ? = post_tag.post_id
	INNER JOIN tag on post_tag.tag_id = tag.id
	`
	if err = db.Select(&tags, sqlStr, pID); err != nil {
		zap.L().Error("there is no tag in db", zap.Int64("pid", pID), zap.Error(err))
		err = nil
	}
	return
}

func GetPostByTagId(id int64, page, size int64) (posts []*models.Post, err error) {
	sqlStr := `
	select distinct p.post_id, p.description, p.title, p.content, p.author_id, p.community_id, p.create_time, p.update_time 
	from post p
	inner join post_tag on p.post_id = post_tag.post_id
	inner join tag on post_tag.tag_id = ?
	order by create_time desc limit ?, ?
	`
	posts = make([]*models.Post, 0, 2)
	if err = db.Select(&posts, sqlStr, id, (page-1)*size, size); err != nil {
		zap.L().Error("there is no post in db", zap.Int64("id", id), zap.Error(err))
		err = nil
	}
	return
}

// UpdateTag 更新数据库标签名称
func UpdateTag(tag *models.Tag) (err error) {
	sqlStr := `update tag set tag_name = ? where id = ?`
	_, err = db.Exec(sqlStr, tag.Name, tag.ID)
	return err
}

// DeleteTagById 根据id删除标签
func DeleteTagById(tid int64) (err error) {
	sqlStr := `delete from tag where id = ?`
	ret, err := db.Exec(sqlStr, tid)
	if err != nil {
		zap.L().Error("there delete is failed", zap.Error(err))
		return ErrorDeleteFailed
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if n == 0 {
		return ErrorTagNotExist
	}
	if err != nil {
		return err
	}
	return
}

// GetTotalByTag 根据tag获取post数量
func GetTotalByTag(tag string) (totalTag int64, err error) {
	sqlStr := `select count(tag_id) from post_tag where tag_id = (
		select id from tag where tag_name = ?
	)`
	err = db.Get(&totalTag, sqlStr, tag)
	return
}

// GetTotalTag 标签总数
func GetTotalTag() (totalTag int64, err error) {
	sqlStr := `select count(id) from tag`
	err = db.Get(&totalTag, sqlStr)
	return
}