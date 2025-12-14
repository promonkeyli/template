package database

import (
	"mall-api/internal/config"
	"mall-api/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() (*gorm.DB, error) {
	dsn := config.NewPostgreDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，必须显示指定，不然gorm默认会使用复数表名
		},
	})
	if err != nil {
		return nil, err
	}

	logger.Log.Info("PostgreSQL 数据库连接成功！")
	return db, nil
}
