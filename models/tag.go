package models

type Tag struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"tag_name" binding:"required"`
}

// ApiPostDetail 标签详情接口的结构体
type ApiTagDetail struct {
	Id   int64          `json:"id"`
	Name string         `json:"name" binding:"required"`
	Post []*ApiPostDetailList `json:"post"`
	TotalPages int64 `json:"total_pages"`
}
