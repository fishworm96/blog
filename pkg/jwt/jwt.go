package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("秘钥")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer:    "blog",                                                                            // 签发人
		},
	}
	// 使用指定的签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// RefreshToken token 续期
func RefreshToken(tokenString, username string, userID int64) (string, error, bool) {
	// 解析原始的 Token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		return "", err, false
	}

	// 检查原始 Token 是否有效
	if !token.Valid {
		return "", errors.New("invalid token"), false
	}

	// 获取原始 Token 中的 claims 信息
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("failed to parse claims"), false
	}

	// 如果Token还有30分钟过期，则生成新的Token
	timeRemaining := time.Until(time.Unix(claims.ExpiresAt, 0))
	if timeRemaining <= 10 * time.Minute {
		newToken, err := GenToken(userID, username)
		if err != nil {
			return "", err, false
		}
		return newToken, nil, true
	}

	// 如果Token未过期且距离过期时间还有超过30分钟，则返回原始Token
	return "", nil, false
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		// 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
