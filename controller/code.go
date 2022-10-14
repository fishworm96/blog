package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserNotExist
	CodeUserExist
	CodeInvalidPassword
	CodeServerBusy
	CodeAddFailed
	CodeUpdateFailed
	CodeDeleteFailed
	CodePostNotExist
	CodeTagExist
	CodeTagNotExist
	CodeNeedLogin
	CodeInvalidToken
	CodeFileSuffixNotLegal
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
	CodeInvalidParam: "请求参数错误",
	CodeUserNotExist: "用户不存在",
	CodeUserExist: "用户名已存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy: "服务繁忙",
	CodeAddFailed: "添加失败",
	CodeUpdateFailed: "修改失败",
	CodeDeleteFailed: "删除失败",
	CodePostNotExist: "帖子不存在",
	CodeTagExist: "标签已存在",
	CodeTagNotExist: "标签不存在",
	CodeNeedLogin: "需要登录",
	CodeInvalidToken: "无效的token",
	CodeFileSuffixNotLegal: "文件后缀名不合法",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}