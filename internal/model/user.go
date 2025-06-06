package model

import (
	"time"
)

type User struct {
	ID          int           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string        `gorm:"type:char(30);not null" json:"name"`
	Account     string        `gorm:"type:char(30);not null;uniqueIndex" json:"account"`
	CreateTime  time.Time     `gorm:"type:datetime(6);" json:"createTime"`
	Password    *string       `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Roles       []*Role       `gorm:"many2many:user_roles"`
	Departments []*Department `gorm:"many2many:user_departments;" json:"departments,omitempty"`
	Phone       string        `gorm:"type:char(30)" json:"phone"`
	Email       string        `gorm:"type:char(100)" json:"email"`
	UpdateTime  time.Time     `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (User) TableName() string {
	return "users"
}
