package config

import (
	"encoding/json"
	"log"
	"os"
)

// DatabaseConfig representa as configurações de conexão do banco de dados.
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

// Config é a estrutura global de configurações.
type Config struct {
	Database DatabaseConfig `json:"database"`
}

// AppConfig será carregada no início da aplicação.
var AppConfig Config

// LoadConfig lê o arquivo de configuração e preenche a variável AppConfig.
func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Não foi possível abrir o arquivo de configuração: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Erro ao decodificar o arquivo de configuração: %v", err)
	}
}
