package model

import (
	"time"
)

type User struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"type:char(30);not null" json:"name"`
	Account    string    `gorm:"type:char(30);not null" json:"account"`
	CreateTime string    `gorm:"type:datetime(6);" json:"create_time"` // 注意：这是 char 类型而不是 time 类型
	Password   string    `gorm:"type:varchar(255);not null" json:"password"`
	Role       string    `gorm:"type:varchar(10)" json:"role"`
	Phone      string    `gorm:"type:char(30)" json:"phone"`
	Email      string    `gorm:"type:char(100)" json:"email"`
	UpdateTime time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`

	// 可选：如果你要关联 Version 模型（1对多）
	// Versions []Version `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (User) TableName() string {
	return "user"
}
