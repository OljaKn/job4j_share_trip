package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	DBName    string
	DBSSLMode string
}

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading env file")
		return
	}
	log.Println("Loading env file")
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		DBName:    os.Getenv("DB_NAME"),
		DBSSLMode: os.Getenv("DB_SSLMODE"),
	}
}

func (c DBConfig) DSN() string {
	ssl := c.DBSSLMode
	if ssl == "" {
		ssl = "disable"
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, ssl,
	)
}
