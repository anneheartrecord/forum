package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

var mySecret = []byte("huangxinci")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Usernaem string `json:"usernaem"`
	jwt.StandardClaims
}

//创建token
func GenToken(userID int64, username string) (string, error) {
	//创建一个我们声明的数据
	c := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer:    "my-project",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//加盐
	return token.SignedString(mySecret)
}

//解析token
func ParseToken(tokenstring string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenstring, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
