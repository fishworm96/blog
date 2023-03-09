package models

import "time"

// 内存对齐
type Post struct {
	ID          string     `json:"id" db:"post_id" swaggerignore:"true"`
	Description string    `json:"description" db:"description" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`     // 文章标题
	Content     string    `json:"content" db:"content" binding:"required"` // 文章内容
	AuthorID    int64     `json:"author_id" db:"author_id" swaggerignore:"true"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区ID
	Tag         []int64   `json:"tag" binding:"required"`
	Status      int32     `json:"status" db:"status" swaggerignore:"true"`
	CreateTime  time.Time `json:"create_time" db:"create_time" swaggerignore:"true"`
	UpdateTime  time.Time `json:"update_time" db:"update_time" swaggerignore:"true"`
}

// ApiPostList 帖子列表接口的结构体
type ApiPostList struct {
	AuthorName       string             `json:"author_name"`
	Tag              []string           `json:"tag"`
	VoteNum          int64              `json:"vote_num"`
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}

// APiPostDetail 帖子接口的结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	Tag              []*Tag           `json:"tag"`
	VoteNum          int64              `json:"vote_num"`
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}