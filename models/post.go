package models

import "time"

// 内存对齐
type Post struct {
	ID int64 `json:"id" db:"post_id" swaggerignore:"true"`
	AuthorID int64 `json:"author_id" db:"author_id" swaggerignore:"true"`
	CommunityID int64 `json:"community_id" db:"community_id" binding:"required"` // 社区ID
	Status int32 `json:"status" db:"status" swaggerignore:"true"`
	Title string `json:"title" db:"title" binding:"required"` // 文章标题
	Content string `json:"content" db:"content" binding:"required"` // 文章内容
	CreateTime time.Time `json:"create_time" db:"create_time" swaggerignore:"true"`
	UpdateTime time.Time `json:"update_time" db:"update_time" swaggerignore:"true"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum int64 `json:"vote_num"`
	*Post // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}