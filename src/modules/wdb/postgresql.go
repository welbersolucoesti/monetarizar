package wdb

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PGConn Conexão Mysql
type pgconn struct{}

// InitDB --
func (s *pgconn) InitDB(config Config) (conn *gorm.DB, err error) {

	var dsn string

	if config.DBScherma != "" {
		dsn = fmt.Sprintf(
			`host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable TimeZone=America/Sao_Paulo`,
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBPassword,
			config.DBName,
			config.DBScherma)
	} else {
		dsn = fmt.Sprintf(
			`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo`,
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBPassword,
			config.DBName)
	}

	conn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Printf("%#v", dsn)
	}

	return
}
