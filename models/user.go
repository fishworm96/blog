package models

type User struct {
	UserID int64 `db:"user_id"`
	IsSuper int64 `db:"is_super"`
	Username string `db:"username"`
	Password string `db:"password"`
}