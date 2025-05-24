package model

import (
	"time"
)

type Roles struct {
	ID          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"type:char(30);not null" json:"name"`
	Users       []*User        `gorm:"many2many:user_roles"`
	Permissions []*Permissions `gorm:"many2many:roles_permissions"`
	Description string         `gorm:"type:char(30);" json:"description"`
	CreateTime  time.Time      `gorm:"type:datetime(6);" json:"create_time"`
	UpdateTime  time.Time      `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`
}

func (Roles) TableName() string {
	return "roles"
}
