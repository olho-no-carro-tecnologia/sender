package database

import (
	"log"

	"poc.sender/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect inicializa a conexão com o banco de dados
func Connect() {
	dsn := config.GetEnv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=meuprojeto port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[ERROR] Falha ao conectar ao banco de dados: %v", err)
	}

	log.Println("[INFO] Conectado ao banco de dados com sucesso")
	DB = db
}

// GetDB retorna a conexão ativa do banco de dados
func GetDB() *gorm.DB {
	return DB
}
