package model

import (
	"time"
)

type RolePermissions struct {
	ID           int         `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleID       int         `gorm:"not null" json:"role_id"`
	Role         Roles       `gorm:"foreignKey:RoleID;not null" json:"role"`
	PermissionID int         `gorm:"not null" json:"permission_id"`
	Permission   Permissions `gorm:"foreignKey:PermissionID;not null" json:"permission"`
	CreateTime   time.Time   `gorm:"type:datetime(6);" json:"create_time"`
	UpdateTime   time.Time   `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`
}

func (RolePermissions) TableName() string {
	return "role_permissions"
}
