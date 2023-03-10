package logic

import (
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/send_email"
	"blog/pkg/snowflake"
	"blog/setting"
	"context"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"

	// "blog/pkg/tools"
	// "blog/setting"
	// "net"
	// "os"
	// "path"
	// "strconv"

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

func EmailLogin(p *models.EmailLogin) (token string, err error) {
	user := new(models.User)
	if err := redis.GetCode(p); err != nil {
		return "", err
	}
	mysql.EmailLogin(user, p.Email)
	return jwt.GenToken(user.UserID, user.Username)
}

func GetUserInfo(uid int64) (data *models.UserInfo, err error) {
	return mysql.GetUserInfoByUserId(uid)
}

func GetUserInfoList() (info []*models.UserInfo, err error) {
	return mysql.GetUserInfoList()
}

func UploadAvatar(file *multipart.FileHeader, extName string, userID int64) (string, error) {
	// 上传到服务器
	// var host string
	// port := ":" + strconv.Itoa(setting.Conf.Port)
	// day := tools.GetDay()
	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	return "", err
	// }
	// for _, address := range addrs {
	// 	// 检查ip地址判断是否回环地址
	// 	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			host = ipnet.IP.String()
	// 		}
	// 	}
	// }
	// dir := "./static/upload" + day
	// err = os.MkdirAll(dir, 0666)
	// if err != nil {
	// 	return "", err
	// }
	// fileName := strconv.FormatInt(tools.GetUnix(), 10) + extName

	// dst := path.Join(dir, fileName)
	// dir2 := host + port + "/" + dst
	// 上传到OSS

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	putPlicy := storage.PutPolicy{
		Scope: setting.Conf.Bucket,
	}
	mac := qbox.NewMac(setting.Conf.AccessKey, setting.Conf.SecretKey)

	// 获取上传凭证
	upToken := putPlicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 华东区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 上传自定义 Key ，可以指定上传目录及文件名和后缀
	key := "avatar/" + file.Filename // 上传路径，如果当前目录中已存在相同文件，则返回失败错误
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	if err != nil {
		return "", err
	}

	url := "http://" + setting.Conf.ImgUrl + ret.Key

	return url, mysql.EditAvatar(url, userID)
}

func SendCode(email string) (err error) {
	code := send_email.GetRand()
	if err = send_email.SendCode(email, code); err != nil {
		zap.L().Error("err", zap.Error(err))
		return
	}
	return redis.SaveCode(email, code)
}
