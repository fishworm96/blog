package tools

import (
	"mime/multipart"
	"path"
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
		".jpg": true,
		".png": true,
		".gif": true,
		".jpeg": true,
	}
	ok = allowExtMap[extName]
	return
}
