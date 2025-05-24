package model

import (
	"time"
)

type Permissions struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:char(30);not null" json:"name"`
	Roles       []*Roles  `gorm:"many2many:roles_permissions"`
	Description string    `gorm:"type:char(30);" json:"description"`
	CreateTime  time.Time `gorm:"type:datetime(6);" json:"create_time"` // 注意：这是 char 类型而不是 time 类型
	UpdateTime  time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`
}

func (Permissions) TableName() string {
	return "permissions"
}
