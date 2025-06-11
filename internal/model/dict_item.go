package model

type DictItem struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	DictCode string `gorm:"index;not null" json:"dictCode"`            // 外键，关联 Dict
	Value    string `gorm:"type:varchar(200);not null" json:"value"`   // 数据值，系统用的
	LabelZh  string `gorm:"type:varchar(200);not null" json:"labelZh"` // 中文名称
	LabelEn  string `gorm:"type:varchar(200)" json:"labelEn"`          // 英文名称
	Sort     int    `gorm:"default:0" json:"sort"`
	Status   string `gorm:"type:char(1);default:'1'" json:"status"` // 启用状态
}

func (DictItem) TableName() string {
	return "dict_items"
}
