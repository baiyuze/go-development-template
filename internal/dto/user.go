package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type LoginBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegBody struct {
	Account  *string `json:"account,omitempty"`
	Password *string `json:"password,omitempty"`
	Name     *string `json:"name,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Email    *string `json:"email,omitempty"`
}

type UserInfo struct {
	Account string  `json:"account"`
	Name    string  `json:"name"`
	Id      float64 `json:"id"`
}

type CustomClaims struct {
	UserID  float64 `json:"sub"`
	Account string  `json:"account"`
	Name    string  `json:"name"`
	jwt.RegisteredClaims
}

type LoginResult struct {
	Token        string    `json:"token,omitempty"`
	RefreshToken string    `json:"refreshToken,omitempty"`
	UserInfo     *UserInfo `json:"userInfo,omitempty"`
}

type UserRoleRequest struct {
	ID      int   `json:"id"`
	RoleIds []int `json:"roleIds"`
}

type User struct {
	Roles []int `json:"roles,omitempty"`
}

type UserWithRole struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Account    string    `json:"account"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	RoleIDs    []int     `json:"roleId"`
	RoleNames  []string  `json:"roleName"`
}
