package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT         string
	DATABASEURL  string
	SMTPHost     string
	SMTPPort     string
	SMTPEmail    string
	SMTPPassword string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		PORT:         os.Getenv("PORT"),
		DATABASEURL:  os.Getenv("DATABASE_URL"),
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPEmail:    os.Getenv("SMTP_EMAIL"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
}
