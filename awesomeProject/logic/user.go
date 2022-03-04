package logic

import (
	"golangstudy/jike/awesomeProject/dao/mysql"
	"golangstudy/jike/awesomeProject/models"
	"golangstudy/jike/awesomeProject/pkg/jwt"
	snowflake "golangstudy/jike/awesomeProject/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	//构造user实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//密码加密

	//保存进数据库
	return mysql.InsertUser(u)
}
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	//生成JWT
	token, err = jwt.GenToken(user.UserID, user.Username)
	return
}
