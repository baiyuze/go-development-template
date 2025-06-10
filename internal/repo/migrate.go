package repo

import (
	"app/internal/model"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&model.User{},
		&model.Department{},
		&model.Role{},
		&model.Permission{},
		&model.Dict{},
		&model.DictItem{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
