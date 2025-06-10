package model

import (
	"time"
)

type Role struct {
	ID          int           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string        `gorm:"type:varchar(30);not null" json:"name"`
	Users       []*User       `gorm:"many2many:user_roles" json:"users"`
	Permissions []*Permission `gorm:"many2many:role_permissions" json:"permissions"`
	Description string        `gorm:"type:varchar(200);" json:"description"`
	CreateTime  time.Time     `gorm:"type:datetime(6);autoCreateTime" json:"createTime"`
	UpdateTime  time.Time     `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (Role) TableName() string {
	return "roles"
}
