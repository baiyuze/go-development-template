package repo

import (
	"log"
	"os"

	"go.uber.org/dig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func InitDB() *Repo {

	// dsn := "test:test123@tcp(192.168.1.1::3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
	SQL_URL := os.Getenv("SQL_URL")
	dsn := SQL_URL
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移模型（可选）
	Migrate(db)
	return &Repo{DB: db}
}

func ProvideDB(container *dig.Container) {
	container.Provide(InitDB)
	container.Provide(func(r *Repo) *gorm.DB {
		return r.DB
	})
}
