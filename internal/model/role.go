package model

import (
	"time"
)

type Role struct {
	ID          int           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string        `gorm:"type:char(30);not null" json:"name"`
	Users       []*User       `gorm:"many2many:user_roles"`
	Permissions []*Permission `gorm:"many2many:role_permissions"`
	Description string        `gorm:"type:char(30);" json:"description"`
	CreateTime  time.Time     `gorm:"type:datetime(6);" json:"create_time"`
	UpdateTime  time.Time     `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`
}

func (Role) TableName() string {
	return "roles"
}
