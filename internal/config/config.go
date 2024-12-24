package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string `env:"SERVER_HOST"`
	Port     string `env:"SERVER_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	DBHost   string `env:"DB_HOST"`
	DBPort   string `env:"DB_PORT"`
	DBName   string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSLMODE"`
}

func (c *Config) FormatDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		c.Username, c.Password, c.DBHost,
		c.DBPort, c.DBName, c.SSLMode)
}
func (c *Config) FormatAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func LoadEnv() Config {
	godotenv.Load()

	return Config{
		Host:     getEnv("SERVER_HOST", "http://localhost"),
		Port:     getEnv("SERVER_PORT", "8881"),
		Username: getEnv("DB_USERNAME", "root"),
		Password: getEnv("DB_PASSWORD", "root"),
		DBHost:   getEnv("DB_HOST", "127.0.0.1"),
		DBPort:   getEnv("DB_PORT", "5432"),
		DBName:   getEnv("DB_NAME", "tempi"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
