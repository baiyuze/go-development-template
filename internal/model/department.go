package model

import (
	"time"
)

type Department struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	ParentID    int       `gorm:"default:0;index" json:"parent_id"`      // 上级部门 ID，0 表示根节点
	Sort        int       `gorm:"default:0" json:"sort"`                 // 排序
	Status      uint8     `gorm:"type:tinyint;default:1;" json:"status"` // 1=启用，0=禁用
	CreatedTime time.Time `gorm:"type:datetime(6);autoCreateTime" json:"created_time"`
	UpdateTime  time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"updated_time"`
	Description string    `gorm:"type:varchar(200);description" json:"description"`
	// 多对多关联
	Users []*User `gorm:"many2many:user_departments;" json:"users,omitempty"`

	// 构造树结构用
	Children []*Department `gorm:"-" json:"children"`
}

func (Department) TableName() string {
	return "departments"
}
