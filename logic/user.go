package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"
	"blog/pkg/tools"
	"blog/setting"
	"mime/multipart"
	"net"
	"os"
	"path"
	"strconv"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
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

func GetUserInfo(uid int64) (data *models.UserInfo, err error) {
	return mysql.GetUserInfoByUserId(uid)
}

func UploadImage(file *multipart.FileHeader, extName string, userID int64) (string, error) {
	var host string
	port := ":" + strconv.Itoa(setting.Conf.Port)
	day := tools.GetDay()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				host = ipnet.IP.String()
			}
		}
	}
	dir := "./static/upload" + day
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		return "", err
	}
	fileName := strconv.FormatInt(tools.GetUnix(), 10) + extName

	dst := path.Join(dir, fileName)
	dir2 := host + port + "/" + dst

	return dst, mysql.EditAvatar(dir2, userID)
}
