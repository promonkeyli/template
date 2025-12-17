package database

import (
	"fmt"
	"time"

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
	// 数据库ssl模式
	SSLMode string

	// 最大空闲连接数
	MaxIdleConns int
	// 最大打开连接数
	MaxOpenConns int
	// 连接最大生存时间
	ConnMaxLifetime int
}

func NewPostgre(c *PostgreConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.TimeZone, c.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，必须显示指定，不然gorm默认会使用复数表名
		},
	})
	if err != nil {
		return nil, err
	}

	// 配置连接池： GORM 使用 database/sql 来维护连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间：配置是以秒为单位
	sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)
	return db, nil
}
