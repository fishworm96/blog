package models

type Tag struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"tag_name" binding:"required"`
}
