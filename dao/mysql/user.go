package mysql

import (
	"blog/models"
	// "crypto/md5"
	"database/sql"
	// "encoding/hex"

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
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
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

// GetUserById 根据用户id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}

func GetRoleHandler(uid int64) (role *models.Role, err error) {
	role = new(models.Role)
	sqlStr := `select role_id, title, description from role where role_id = (
		select role_id from user where user_id = ?
	)`
	err = db.Get(role, sqlStr, uid)
	return
}