package models

import (
	"time"
)

type Version struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:char(30);not null" json:"name"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`                 // 外键字段
	Status      uint8     `gorm:"type:tinyint;default:0;not null" json:"status"` // 0: 初始化, 1: 成功, 2: 失败
	GitUrl      string    `gorm:"type:text" json:"git_url"`
	Description string    `gorm:"type:char(255)" json:"description"`
	CreateTime  time.Time `gorm:"type:datetime(6);autoCreateTime" json:"create_time"`
	UpdateTime  time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`

	// 可选：关联的用户（需要定义 User 模型）
	User User `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

// 表名（如果不想让 GORM 自动转复数）
func (Version) TableName() string {
	return "version"
}
