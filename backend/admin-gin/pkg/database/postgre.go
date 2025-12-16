package database

import (
	"fmt"
	"mall-api/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgreConfig struct {
	// 主机ip地址
	Host string
	// 数据库登录用户名
	User string
	// 数据库登录密码
	Password string
	// 数据库名称
	DBName string
	// 数据库端口
	Port int
	// 数据库时区
	TimeZone string
}

func NewPostgre(c *PostgreConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.TimeZone,
	)
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
