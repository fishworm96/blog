package models

type MenuDetail struct {
	Id int64 `json:"id" db:"id"`
	ModuleId int64 `json:"module_id" db:"module_id"`
	Title string `json:"title" db:"title"`
	Icon string `json:"icon" db:"icon"`
	Path string `json:"path" db:"path"`
	Children []*MenuDetail `json:"children,omitempty"` // 过滤为空
}
