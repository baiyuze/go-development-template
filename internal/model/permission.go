package model

import (
	"time"
)

type Permission struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:char(30);not null" json:"name"`
	Roles       []*Role   `gorm:"many2many:role_permissions"`
	Description string    `gorm:"type:char(30);" json:"description"`
	CreateTime  time.Time `gorm:"type:datetime(6);" json:"createTime"`
	UpdateTime  time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (Permission) TableName() string {
	return "permissions"
}
