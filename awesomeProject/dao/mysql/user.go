package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"golangstudy/jike/awesomeProject/models"
)

const secret = "lovehuangxinci"

func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encyptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}
func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}
func encyptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登录的密码
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err != nil {
		return ErrorUserNotExist
	}
	//判断密码
	password := encyptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id=?`
	err = db.Get(user, sqlStr, uid)
	return
}
