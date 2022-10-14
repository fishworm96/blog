package models

type Role struct {
	Role_id int64 `json:"role_id" db:"role_id"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type RoleInfo struct {
	Username string `json:"username"`
	Title string `json:"title"`
	Description string `json:"description"`
}