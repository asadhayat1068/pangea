package config

import (
	"fmt"
	"os"
)

// DB holds PostgreSQL connection settings, loaded from the environment.
type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// LoadDB reads settings from DB_* env vars, with local-dev defaults
// that match docker-compose.yml

func LoadDB() DB {
	return DB{
		Host:     getenv("DB_HOST", "localhost"),
		Port:     getenv("DB_PORT", "54321"),
		User:     getenv("DB_USER", "pangea"),
		Password: getenv("DB_PASSWORD", "pangea"),
		Name:     getenv("DB_NAME", "pangea"),
		SSLMode:  getenv("DB_SSLMODE", "disable"),
	}
}

// DSN returns a libpq key=value connection string (used by database/sql later).
func (d DB) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}

// URL returns a URL-style connection string (used by golang-migrate).
func (d DB) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode)
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
