package mysql

import "errors"

var (
	ErrorGetFailed = errors.New("获取失败")
	ErrorUserExist = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户或密码错误")
	ErrorInvalidID = errors.New("无效的ID")
	ErrorAddFailed = errors.New("添加失败")
	ErrorUpdateFailed = errors.New("更新失败")
	ErrorDeleteFailed = errors.New("删除失败")
	ErrorImageUrFailed = errors.New("上传失败")
	ErrorPostNotExist = errors.New("帖子不存在")
	ErrorTagExist = errors.New("标签已存在")
	ErrorTagNotExist = errors.New("标签不存在")
	ErrorMenuNotExist = errors.New("菜单不存在")
)