package config

import (
	"log"
	"os"

	"gopkg.in/go-playground/validator.v9"
)

// Config contains project configuration
type Config struct {
	Port string `validate:"required"`

	DBHost     string `validate:"required"`
	DBUserName string `validate:"required"`
	DBPassword string `validate:"required"`
	DBPort     string `validate:"required"`
	DBName     string `validate:"required"`

	JwtSigningKey string `validate:"required"`
}

var config Config

func init() {
	config.Port = ":" + os.Getenv("PORT")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_NAME")
	config.JwtSigningKey = os.Getenv("JWT_SIGNING_KEY")

	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		log.Fatalf("Environment Variables not set:\n%s\n", err)
	}
}

// Get configuration structure
func Get() Config {
	return config
}
