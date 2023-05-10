package mysql

import (
	"blog/models"
	// "crypto/md5"
	"database/sql"
	// "encoding/hex"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// const secret = "12345"

// CheckUserExist 检查指定用户的用户名是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 执行sql语句入库
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password, is_super, role_id, nick_name) values(?, ?, ?, 0, 2, "user")`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(oPassword string) string {
	password := []byte(oPassword)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashedPassword)
	// h := md5.New()
	// h.Write([]byte(secret))
	// return hex.EncodeToString(h.Sum([]byte(oPassword)))

}

// Login 查询数据库判断账号和密码是否正确
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	// password := encryptPassword(oPassword)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oPassword))
	if err != nil {
		return ErrorInvalidPassword
	}
	// if password != user.Password {
	// 	return ErrorInvalidPassword
	// }
	return
}

func EmailLogin(user *models.User, email string) (err error) {
	sqlStr := `select user_id, username from user where email = ?`
	err = db.Get(&user, sqlStr, email)
	if err != nil {
		return err
	}
	return
}

// GetUserById 根据用户id获取用户信息
func GetUserById(uid int64) (user *models.ArticleAuthor, err error) {
	user = new(models.ArticleAuthor)
	sqlStr := `select user_id, nick_name from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	if err == sql.ErrNoRows {
		err = ErrorUserNotExist
		return nil, err
	}
	return
}

func EditAvatar(dst string, userId int64) error {
	sqlStr := `update user set avatar = ? where user_id = ?`
	ret, err := db.Exec(sqlStr, dst, userId)
	if err != nil {
		zap.L().Error("update failed", zap.Error(err))
		return ErrorUpdateFailed
	}
	n, err := ret.RowsAffected()
	if n == 0 {
		return ErrorUserNotExist
	}
	return err
}

func GetUserInfoByUserId(uid int64) (info *models.UserInfo, err error) {
	info = new(models.UserInfo)
	sqlStr := `select username, email, nick_name, avatar, is_super, role_id, gender from user where user_id = ?`
	if err = db.Get(info, sqlStr, uid); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorUserNotExist
			return
		}
	}
	return
}

func GetUserInfoList() (info []*models.UserInfo, err error) {
	sqlStr := `select username, email, nick_name, avatar, is_super, role_id, gender from user`
	if err = db.Select(&info, sqlStr); err != nil {
		zap.L().Error("there is no userInfo in db")
		err = nil
	}
	return
}