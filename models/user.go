package models

type User struct {
	UserID int64 `db:"user_id"`
	IsSuper int64 `db:"is_super"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type Role struct {
	Role_id int64 `json:"role_id" db:"role_id"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type UserInfo struct {
	Username string `json:"username"`
	Title string `json:"title"`
	Description string `json:"description"`
}