package models

// 定义请求的参数结构体
const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required,max=24,min=6"`        // 用户名
	Password   string `json:"password" binding:"required,max=24,min=6"`        // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 重复密码
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

type EmailLogin struct {
	Email string `json:"email" binding:"required"` // 邮箱
	Code  string `json:"code" binding:"required"`  // 验证码
}

// ParamPost
type ParamPost struct {
	Title   string `json:"title" binding:"required"`   // 文章标题
	Content string `json:"content" binding:"required"` // 文章内容
}

// 帖子id和点赞两个参数
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 增长票1 还是返回票-1 取消票0
}

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"1"`       // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

type ParamPostAndTag struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ParamMenu struct {
	Title    string `json:"title" db:"title" binding:"required"`
	Icon     string `json:"icon" db:"icon"`
	Path     string `json:"path" db:"path" binding:"required"`
	Type     *int64  `json:"type" db:"type" binding:"required"`
	ModuleID *int64  `json:"module_id" db:"module_id" binding:"required"`
}

type ParamUpdateMenu struct {
	ID       int64  `json:"id" db:"id" binding:"required"`
	Type     *int64  `json:"type" db:"type" binding:"required"`
	ModuleID *int64  `json:"module_id" db:"module_id" binding:"required"`
	Title    string `json:"title" db:"title" binding:"required"`
	Icon     string `json:"icon" db:"icon"`
	Path     string `json:"path" db:"path" binding:"required"`
}

type RoleMenu struct {
	RoleID   int64   `json:"role_id" db:"role_id" binding:"required"`
	AccessID []int64 `json:"access_id" db:"access_id" binding:"required"`
}
