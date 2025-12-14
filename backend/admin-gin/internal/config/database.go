package config

import "fmt"

/**
 * @description PostgreSQL data source name (DSN)
 * @return string
 */
func NewPostgreDSN() string {
	conf := NewServerLocalConfig()

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s",
		conf.Host, conf.User, conf.Password, conf.DBName, conf.Port, conf.TimeZone,
	)
}
