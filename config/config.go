package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT          string
	DATABASEURL   string
	SMTPEmail     string
	ResendAPIKey  string
}

var AppConfig Config

func LoadConfig() {
	_ = godotenv.Load()

	AppConfig = Config{
		PORT:         os.Getenv("PORT"),
		DATABASEURL:  os.Getenv("DATABASE_URL"),
		SMTPEmail:    os.Getenv("SMTP_EMAIL"),
		ResendAPIKey: os.Getenv("RESEND_API_KEY"),
	}
}
