package models

type User struct {
	UserID int64 `db:"user_id"`
	IsSuper int64 `db:"is_super"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type ArticleAuthor struct {
	UserID int64 `json:"user_id" db:"user_id"`
	NickName string `json:"nick_name" db:"nick_name"`
}

type UserInfo struct {
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	NickName string `json:"nick_name" db:"nick_name"`
	Avatar string `json:"avatar" db:"avatar"`
	IsSuper int64 `json:"is_super" db:"is_super"`
	RoleId int64 `json:"role_id" db:"role_id"`
	Gender int64 `json:"gender" db:"gender"`
}

type Email struct {
	Email string `json:"email" binding:"required"`
}