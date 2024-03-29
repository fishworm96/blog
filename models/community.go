package models

import (
	"time"
)

type Community struct {
	ID    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Image string `json:"image" db:"image" binding:"required"`
}

type CommunityDetail struct {
	Community
	Introduction string    `json:"introduction,omitempty" db:"introduction"` // 过滤为空
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}

type CommunityPost struct {
	CommunityDetail   *CommunityDetail     `json:"detail"`
	ApiPostDetailList []*ApiPostDetailList `json:"post_list"`
	TotalPages        int64                `json:"total_pages"`
}

type CommunityPostList struct {
	ID    int64   `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Image string  `json:"image"`
	Post  []*Post `json:"post"`
}

type CommunityCreateDetail struct {
	Name         string `form:"name" db:"name" binding:"required"`
	Introduction string `form:"introduction,omitempty" json:"introduction,omitempty" db:"introduction"` // 过滤为空
	Image        string `form:"image" db:"image" binding:"required"`
}
