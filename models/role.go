package models

type Role struct {
	RoleID int64 `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
}

type RoleInfo struct {
	Username string `json:"username"`
	Title string `json:"title"`
	Description string `json:"description"`
}