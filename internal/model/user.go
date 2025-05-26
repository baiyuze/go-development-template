package model

import (
	"time"
)

type User struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"type:char(30);not null" json:"name"`
	Account    string    `gorm:"type:char(30);not null" json:"account"`
	CreateTime time.Time `gorm:"type:datetime(6);" json:"create_time"`
	Password   *string   `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Roles      []*Role   `gorm:"many2many:user_roles"`
	Phone      string    `gorm:"type:char(30)" json:"phone"`
	Email      string    `gorm:"type:char(100)" json:"email"`
	UpdateTime time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"update_time"`
}

func (User) TableName() string {
	return "users"
}
