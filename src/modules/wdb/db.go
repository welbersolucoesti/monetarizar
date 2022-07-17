package wdb

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Config  Estrutura de confirurações Genéricas
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBScherma  string
}

var db []*gorm.DB

// InitDB Inicia a conexão com o BD
func InitDB(driver string, config Config) {

	var conn *gorm.DB
	var err error

	if driver == "postgresql" {
		var pgconn pgconn
		conn, err = pgconn.InitDB(config)
	}

	if driver == "mysql" {
		var myconn mysqlconn
		conn, err = myconn.InitDB(config)
	}

	if err == nil {
		db = append(db, conn)
	} else {
		log.Fatal(err.Error())
	}
}

// GetConn Retorna uma Conexão do Banco de Dados
func GetConn(id int64) *gorm.DB {
	return db[id]
}

// GetDB Retorna uma Conexão dentro de um escopo
func GetDB(driver string, config Config) (dbConn *gorm.DB, err error) {

	if driver == "postgresql" {
		var pgconn pgconn
		return pgconn.InitDB(config)
	}

	if driver == "mysql" {
		var myconn mysqlconn
		return myconn.InitDB(config)
	}

	return nil, fmt.Errorf("Erro na operação. Driver não encontrado")
}

// Migrate Efetua a migração de todas as tabelas do Banco de Dados
func Migrate(conn *gorm.DB, table interface{}) {
	_ = conn.AutoMigrate(table)
}
