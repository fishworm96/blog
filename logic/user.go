package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"

	"go.uber.org/zap"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 返回token
	return jwt.GenToken(user.UserID, user.Username)
}

func GetUserInfoByUserIdHandler(uid int64) (data *models.UserInfo, err error) {
	user, err := mysql.GetUserById(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
	}
	role, err := mysql.GetRoleHandler(uid)
	if err != nil {
		zap.L().Error("mysql.GetUserById(uid), failed", zap.Int64("uid", uid))
	}
	data = &models.UserInfo{
		Username: user.Username,
		Title: role.Title,
		Description: role.Description,
	}
	return
}