package models

// 定义请求的参数结构体
const (
	OrderTime = "time"
	OrderScore = "score"
)

type ParamSignUp struct{
	Username string `json:"username" binding:"required,max=24,min=6"` // 用户名
	Password string `json:"password" binding:"required,max=24,min=6"` // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 重复密码
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// ParamPost
type ParamPost struct {
	Title string `json:"title" binding:"required"` // 文章标题
	Content string `json:"content" binding:"required"` // 文章内容
}

// 帖子id和点赞两个参数
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID string `json:"post_id" binding:"required"` // 帖子id
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"` // 增长票1 还是返回票-1 取消票0
}

type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"` // 可以为空
	Page int64 `json:"page" form:"page" example:"1"` // 页码
	Size int64 `json:"size" form:"size" example:"1"` // 每页数据量
	Order string `json:"order" form:"order" example:"score"` // 排序依据
}