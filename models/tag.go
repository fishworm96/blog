package models

type Tag struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"tag_name" binding:"required"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiTagDetail struct {
	Id   int64  `json:"id"`
	Name string `json:"name" binding:"required"`
	Post []*ApiPostDetail `json:"post"`
}