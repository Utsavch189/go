package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetEnv(key string, fallback ...string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	defaultValue := ""
	if len(fallback) > 0 {
		defaultValue = fallback[0]
	}
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_HOST"), GetEnv("DB_PORT"), GetEnv("DB_SCHEMA"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening MySQL connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging MySQL: %v", err)
	}

	return db, nil
}
