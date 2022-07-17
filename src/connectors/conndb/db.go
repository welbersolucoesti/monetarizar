package conndb

import (
	"monetizar/src/modules/wdb"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// InitDB Inicia a conexão com o BD
func InitDB() {

	godotenv.Load("cfg.env")

	wdb.InitDB("postgresql", wdb.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	})
}

// GetConn Retorna uma Conexão do Banco de Dados
func GetConn(id int64) *gorm.DB {
	return wdb.GetConn(id)
}
