package dto

import "github.com/golang-jwt/jwt/v5"

type LoginBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserInfo struct {
	Account string
	Name    string
	Id      float64
}

type CustomClaims struct {
	UserID  float64 `json:"sub"`
	Account string  `json:"account"`
	Name    string  `json:"name"`
	jwt.RegisteredClaims
}

type LoginResult struct {
	Token      string    `json:"token,omitempty"`
	FlushToken string    `json:"flushToken,omitempty"`
	UserInfo   *UserInfo `json:"userInfo,omitempty"`
}
