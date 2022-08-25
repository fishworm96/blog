package models

type ParamSignUp struct{
	Username string `json:"username" binding:"required,max=24,min=6"`
	Password string `json:"password" binding:"required,max=24,min=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamPost
type ParamPost struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	PostID int64 `json:"post_id" binding:"required"`
}