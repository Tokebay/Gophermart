package config

import (
	"flag"
	"os"
)

// Config структура для хранения конфигураций сервиса.
type Config struct {
	RunAddress           string
	DatabaseURI          string
	AccrualSystemAddress string
}

// NewConfig создает новый экземпляр структуры Config, читая значения из переменных окружения и флагов командной строки.
func NewConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.RunAddress, "a", "", "Address and port to run the service")
	flag.StringVar(&config.DatabaseURI, "d", "", "Database connection URI")
	flag.StringVar(&config.AccrualSystemAddress, "r", "", "Accrual system address")

	flag.Parse()

	config.parseEnv()

	return config
}

// parseEnv читает значения конфигурации из переменных окружения.
func (c *Config) parseEnv() {
	if envRunAddress := os.Getenv("RUN_ADDRESS"); envRunAddress != "" {
		c.RunAddress = envRunAddress
	}

	if envDBURI := os.Getenv("DATABASE_URI"); envDBURI != "" {
		c.DatabaseURI = envDBURI
	}

	if envAccrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrualSystemAddress != "" {
		c.AccrualSystemAddress = envAccrualSystemAddress
	}
}
