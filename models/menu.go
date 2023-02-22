package models

type MenuDetail struct {
	ID       int64         `json:"id" db:"id"`
	ModuleID int64         `json:"module_id" db:"module_id"`
	Title    string        `json:"title" db:"title"`
	Icon     string        `json:"icon" db:"icon"`
	Path     string        `json:"path" db:"path"`
	Children []*MenuDetail `json:"children,omitempty"` // 过滤为空
}

type MenuDetailInfo struct {
	ID         int64  `json:"id" db:"id"`
	ModuleID   int64  `json:"module_id" db:"module_id"`
	ParentID   int64  `json:"parent_id" db:"parent_id"`
	Title      string `json:"title" db:"title"`
	Icon       string `json:"icon" db:"icon"`
	Path       string `json:"path" db:"path"`
	ParentName string `json:"parent_name" db:"parent_name"`
}
