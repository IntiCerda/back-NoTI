package configs

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort  int
	Environment string
	Debug       bool
}

func GetConfig() *Config {
	config := &Config{
		ServerPort:  8080,
		Environment: "development",
		Debug:       true,
	}

	// Sobreescribir con variables de entorno si existen (????????????)
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.ServerPort = p
		}
	}

	if env := os.Getenv("ENVIRONMENT"); env != "" {
		config.Environment = env
	}

	if debug := os.Getenv("DEBUG"); debug == "false" {
		config.Debug = false
	}

	return config
}
