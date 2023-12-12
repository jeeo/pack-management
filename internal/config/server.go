package config

import (
	"log"
	"os"
	"time"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DBName string `env:"DB_NAME"`
	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT"`
	DBUser string `env:"DB_USER"`
	DBPwd  string `env:"DB_PASSWORD"`
}

type ServerConfig struct {
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"30s"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"30s"`
	Port         int           `env:"PORT" envDefault:"3000"`

	DBConfig DatabaseConfig
}

func LoadConfig() ServerConfig {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "DEV"
	}

	if appEnv == "DEV" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	config := ServerConfig{}
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
