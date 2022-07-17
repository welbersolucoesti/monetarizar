package wdb

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MysqlConn Conexão Mysql
type mysqlconn struct{}

// InitDB --
func (s *mysqlconn) InitDB(config Config) (conn *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)

	conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Printf("%#v", dsn)
	}

	return
}
