package tools

import (
	"blog/models"
	"mime/multipart"
	"path"
	"sort"
	"time"
)

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func SuffixName(file *multipart.FileHeader) (extName string, ok bool) {
	extName = path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	ok = allowExtMap[extName]
	return
}

// 将一维数组转成树形结构
func GetTreeRecursive(list []*models.MenuDetail, parentId int64) []*models.MenuDetail {
	res := make([]*models.MenuDetail, 0)
	for _, v := range list {
		if v.ModuleID == parentId {
			v.Children = GetTreeRecursive(list, v.ID)
			res = append(res, v)
		}
	}

	// Sort res by ID
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})

	return res
}
