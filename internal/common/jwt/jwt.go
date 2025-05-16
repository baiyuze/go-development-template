package jwt

import (
	"app/internal/dto"
	"app/internal/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Auth 认证jwt，返回token
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

// Analysis 解析token
func Analysis(tokenString string) (dto.UserInfo, error) {
	claims := &dto.CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		appName := "app"
		return []byte(appName), nil
	})

	switch {
	case token.Valid:
		return dto.UserInfo{
			Account: claims.Account,
			Name:    claims.Name,
			Id:      claims.UserID,
		}, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		return dto.UserInfo{}, jwt.ErrTokenMalformed
	case errors.Is(err, jwt.ErrTokenMalformed):
		return dto.UserInfo{}, errors.New("token signature is invalid")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return dto.UserInfo{}, errors.New("token expired")
	default:
		return dto.UserInfo{}, err
	}
}

// 校验token是否有效
func VerifyValidByToken() {

}
