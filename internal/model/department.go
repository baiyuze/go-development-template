package model

import (
	"time"
)

type Department struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	ParentID  uint64    `gorm:"default:0;index" json:"parent_id"`     // 上级部门 ID，0 表示根节点
	Sort      int       `gorm:"default:0" json:"sort"`                // 排序
	Status    uint8     `gorm:"type:tinyint;default:1" json:"status"` // 1=启用，0=禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 多对多关联
	Users []*User `gorm:"many2many:user_departments;" json:"users,omitempty"`

	// 构造树结构用
	Children []*Department `gorm:"-" json:"children,omitempty"`
}

func (Department) TableName() string {
	return "departments"
}
