package database

import (
	"context"
	"log"

	"agnikriti_admin_backend/config"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	conn, err := pgx.Connect(
		context.Background(),
		config.AppConfig.DATABASEURL,
	)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	DB = conn

	log.Println("Database connected successfully")
}
