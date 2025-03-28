package config

import (
	"fmt"
	"os"
	"strconv"
)

// DatabaseConfig tiene la configuración específica de base de datos
type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type ServerConfig struct {
	Port int
}

// NewConfig crea una nueva instancia de configuración con validaciones
func NewConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
	if err != nil {
		return nil, fmt.Errorf("invalid database port: %v", err)
	}

	serverPort, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid server port: %v", err)
	}

	config := &Config{
		Database: DatabaseConfig{
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			Name:     getEnv("DB_NAME", "transport_challenge"),
		},
		Server: ServerConfig{
			Port: serverPort,
		},
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) validate() error {
	if c.Database.User == "" {
		return fmt.Errorf("database user is required")
	}

	if c.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}

	if c.Database.Port <= 0 {
		return fmt.Errorf("invalid database port")
	}

	if c.Server.Port <= 0 {
		return fmt.Errorf("invalid server port")
	}

	return nil
}
