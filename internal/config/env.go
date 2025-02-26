package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente do arquivo .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] Nenhum arquivo .env encontrado, usando variáveis de ambiente do sistema")
	}
}

// GetEnv retorna o valor de uma variável de ambiente com um fallback opcional
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
