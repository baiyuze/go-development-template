package model

import (
	"time"
)

type Dict struct {
	ID          int         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string      `gorm:"type:varchar(100);not null" json:"name"`
	En          string      `gorm:"type:varchar(50)" json:"en"`
	Code        string      `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
	Description string      `gorm:"type:varchar(200)" json:"description"`
	CreateTime  time.Time   `gorm:"type:datetime(6);autoUpdateTime;index" json:"createTime"`
	UpdateTime  time.Time   `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
	Items       []*DictItem `gorm:"foreignKey:DictCode;references:Code;" json:"items"`
}

func (Dict) TableName() string {
	return "dicts"
}
