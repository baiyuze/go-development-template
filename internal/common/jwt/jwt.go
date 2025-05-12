package jwt

import (
	"app/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 认证jwt，返回token
func Auth(user model.User, exp int64) (string, error) {
	appName := "app"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"name":    user.Name,
		"account": user.Account,
		"iat":     time.Now().Unix(),
		"exp":     exp,
		"nbf":     time.Now().Unix(),
	})
	sign, err := token.SignedString([]byte(appName))
	if err != nil {
		return "", err
	}
	return sign, nil
}

// 解析token
func Analysis() {

}

// 校验token是否有效
func VerifyValidByToken() {

}
