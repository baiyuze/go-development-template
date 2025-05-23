package model

import (
	"time"
)

type UserRoles struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:char(30);not null" json:"name"`
	RoleID      int       `gorm:"not null" json:"role_id"`
	Role        Roles     `gorm:"foreignKey:RoleID" json:"role"`
	UserID      int       `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;not null" json:"user"`
	Description string    `gorm:"type:char(30);" json:"description"`
	CreateTime  time.Time `gorm:"type:datetime(6);" json:"create_time"` // 注意：这是 char 类型而不是 time 类型
	UpdateTime  time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`
}

func (UserRoles) TableName() string {
	return "user_roles"
}
