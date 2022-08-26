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
}

// 帖子id和点赞两个参数
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID string `json:"post_id" binding:"required"` // 帖子id
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"` // 增长票1 还是返回票-1 取消票0

}